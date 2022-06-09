package main

import (
	service "billing/cmd/service"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pkg/config"
	"pkg/gateway"
	"pkg/tracing"
	"syscall"
)

const (
	zipkinURL   = "http://localhost:9411/api/v2/spans"
	serviceName = "billing-gateway"
)

func RunGateway() {
	var (
		consulHost = flag.String("consul.host", "127.0.0.1", "consul server ip address")
		consulPort = flag.Int("consul.port", 8500, "consul server port")
	)
	logger := config.GetKitLogger("ReverseProxy")

	logger.Log("tracer", "Zipkin", "URL", zipkinURL)
	tracingImpl, err := tracing.NewOpenTracingTracer(serviceName)
	if err != nil {
		logger.Log("new zipkin tracer", "failed")
		os.Exit(-1)
	}
	tracer := tracingImpl
	defer tracingImpl.Reporter.Close()

	proxy, err := gateway.NewReverseProxy(*consulHost, *consulPort, logger)
	if err != nil {
		logger.Log("err", err)
		os.Exit(-1)
	}

	consumer, err := gateway.InitQueueServer(5, "billing_queue")
	if err != nil {
		logger.Log("method", "InitQueueServer", "err", err)
		os.Exit(-1)
	}
	defer consumer.Conn.Close()
	gateway.ProxySendReq(logger, consumer, tracer.NativeTracer, "http://localhost:10010")
	consumer.Consume(serviceName, 5)

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("ReverseProxy", "Listening", "addr", "10010")
		errc <- http.ListenAndServe(":10010", proxy)
	}()

	logger.Log("exit", <-errc)
}

func main() {
	go RunGateway()

	service.Run()
}
