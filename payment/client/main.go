package main

import (
	"context"
	"fmt"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
	grpc1 "payment/client/grpc"
	http1 "payment/client/http"
	"pkg/discover"
	"pkg/tracing"
)

func main() {
	{	//discover
		discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
		if err != nil {
			log.Printf("1, err: %v\n", err)
			return
		}
		instances, err := discoverclient.DiscoverServices("Payment", "", true)
		if err != nil {
			log.Printf("2, err: %v\n", err)
			return
		}
		fmt.Println(discover.GetInstance(instances))
	}
	{	//http
		reporter := http.NewReporter(tracing.DefaultZipkinURL) //zipkin
		defer reporter.Close()
		zkTracer, err := zipkingo.NewTracer(reporter)
		if err != nil {
			log.Printf("New HTTP TracingImpl failed, err: %v", err)
			return
		}
		zkClientTrace := kitzipkin.HTTPClientTrace(zkTracer)
		conn, err := http1.New("127.0.0.1:8081", map[string][]kithttp.ClientOption{
			"Pay" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Pay")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		for i := 0; i < 100; i++ {
			res, err := conn.Pay(ctx, 9999, 9999, "253rfe64tgrw")
			if err != nil {
				log.Printf("conn Pay failed, err: %v", err)
				return
			}
			fmt.Println(res)
		}
	}
	{ //grpc
		reporter := http.NewReporter(tracing.DefaultZipkinURL)
		defer reporter.Close()
		zkTracer, err := zipkingo.NewTracer(reporter)
		if err != nil {
			log.Printf("New GRPC TracingImpl failed, err: %v", err)
			return
		}
		zkClientTrace := kitzipkin.GRPCClientTrace(zkTracer)
		conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
		if err != nil {
			log.Println("1", err)
			return
		}
		defer conn.Close()
		svc2, err := grpc1.New(conn, map[string][]grpc2.ClientOption{
			"Pay": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Pay")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		for i := 0; i < 100; i++ {
			r2, _ := svc2.Pay(ctx, 3124, 1514, "34rfey345re")
			fmt.Println(r2)
		}
	}
}
