package service

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	"log"
	"net"
	http2 "net/http"
	"notification/pkg/config"
	endpoint "notification/pkg/endpoint"
	grpc "notification/pkg/grpc"
	http1 "notification/pkg/http"
	service "notification/pkg/service"
	"os"
	"os/signal"
	"pkg/discover"
	"pkg/pb"
	"pkg/promtheus"
	"pkg/tracing"
	"strconv"
	"strings"
	"syscall"

	"google.golang.org/grpc/health/grpc_health_v1"

	endpoint1 "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"
)

var tracer *tracing.TracingImpl
var logger kitlog.Logger

var fs = flag.NewFlagSet("notification", flag.ExitOnError)
var debugAddr = fs.String("debug-addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var zipkinURL = fs.String("zipkin-url", tracing.DefaultZipkinURL, "Enable Zipkin tracing via a collector URL e.g. http://localhost:9411/api/v1/spans")
var serviceName = fs.String("service-name", "Notification", "default service name")
var consulAddr = fs.String("consul-addr", "127.0.0.1", "consul listen addr")
var consulPort = fs.Int("consul-port", 8500, "consul list port")

func Run() {
	fs.Parse(os.Args[1:])

	logger = config.GetKitLogger()

	if *zipkinURL != "" {
		logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
		tracingImpl, err := tracing.NewOpenTracingTracer(*serviceName)
		if err != nil {
			logger.Log("new zipkin tracer", "failed")
			os.Exit(-1)
		}
		tracer = tracingImpl
		defer tracingImpl.Reporter.Close()
	} else {
		logger.Log("tracer", "none")
		tracer.Tracer = opentracinggo.GlobalTracer()
	}
	/////////////////////////////////////
	discoverClient, err := discover.NewDiscoverClient(*consulAddr, *consulPort, true)
	if err != nil {
		logger.Log("NewDiscoverClient failed", err)
	}
	ss := strings.Split(*grpcAddr, ":")
	num, _ := strconv.Atoi(ss[1])
	instanceID, ok := discoverClient.Register(*serviceName, "", "127.0.0.1", num, nil, logger)
	defer discoverClient.DeRegister(instanceID, logger)
	if !ok {
		log.Printf("service %s register failed", *serviceName)
		os.Exit(-1)
	}
	//////////////////////////////////////////////////////////
	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)

	httpHandler := http1.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}
func getServiceMiddleware(logger kitlog.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{
		service.LoggingMiddleware(logger),
	}

	return
}
func getEndpointMiddleware(logger kitlog.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{
		"NoticeTrip": {
			endpoint.LoggingMiddleware(logger),
			endpoint.InstrumentingMiddleware(promtheus.NewHistogram(config.System, config.MethodNoticeTrip, "NoticeTrip histogram")),
			endpoint.CountingMiddleware(promtheus.NewCounter(config.System, config.MethodNoticeTrip, "NoticeTrip count")),
			zipkin.TraceEndpoint(tracer.NativeTracer, config.MethodNoticeTrip + "_zipkin"),

		},
		"NoticeBill": {
			endpoint.LoggingMiddleware(logger),
			endpoint.InstrumentingMiddleware(promtheus.NewHistogram(config.System, config.MethodNoticeBill, "NoticeBill histogram")),
			endpoint.CountingMiddleware(promtheus.NewCounter(config.System, config.MethodNoticeBill, "NoticeBill count")),
			zipkin.TraceEndpoint(tracer.NativeTracer, config.MethodNoticeBill + "_zipkin"),

		},
	}

	return
}
func initMetricsEndpoint(g *group.Group) {
	http2.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http2.Serve(debugListener, http2.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer(grpc1.UnaryInterceptor(grpc2.Interceptor))
		pb.RegisterNotificationServer(baseServer, grpcServer)
		grpc_health_v1.RegisterHealthServer(baseServer, &discover.HealthImpl{})
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

}
