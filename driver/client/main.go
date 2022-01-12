package main

import (
	"context"
	"driver"
	grpc1 "driver/client/grpc"
	http1 "driver/client/http"
	"fmt"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
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
		instances, err := discoverclient.DiscoverServices("Driver", "", true)
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
			"Driver" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Driver")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &driver.DriverInfoRequest{
			Username: "aaaaaaaaaaaa",
			Password: "bbbbbbbbbb",
		}
		for i := 0; i < 100; i++ {
			childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			res, err := conn.GetDriverInfo(ctx, data)
			if err != nil {
				log.Printf("conn Driver failed, err: %v", err)
				return
			}
			fmt.Println(res)
			childSpan.Finish()
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
			"Driver": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Driver")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &driver.TakeOrderRequest{
			DriverId:   1234556,
			DriverName: "王老五",
			Location:   "街道口",
			Car:        "奔驰",
		}
		for i := 0; i < 100; i++ {
			childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			r2, err := svc2.TakeOrder(ctx, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(r2)
			childSpan.Finish()
		}
	}
}
