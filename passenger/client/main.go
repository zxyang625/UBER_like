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
	"passenger"
	grpc1 "passenger/client/grpc"
	http1 "passenger/client/http"
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
		instances, err := discoverclient.DiscoverServices("Passenger", "", true)
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
			"GetPassengerInfo" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Passenger")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &passenger.GetPassengerInfoRequest{
			Username: "124123541",
			Password: "235tewrfes",
		}
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			res, err := conn.GetPassengerInfo(ctx1, data)
			if err != nil {
				log.Printf("conn Passenger failed, err: %v", err)
				return
			}
			fmt.Println(res)
			//childSpan.Finish()
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
			"PublishOrder": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Passenger")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &passenger.PublishOrderRequest{
			PassengerId:   3124,
			StartTime:     052,
			Origin:        "",
			Destination:   "",
			PassengerName: "",
		}
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			r2, err := svc2.PublishOrder(ctx1, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(r2)
			//childSpan.Finish()
		}
	}
}
