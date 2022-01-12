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
	"pkg/discover"
	"pkg/tracing"
	"trip"
	grpc1 "trip/client/grpc"
	http1 "trip/client/http"
)

func main() {
	{	//discover
		discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
		if err != nil {
			log.Printf("1, err: %v\n", err)
			return
		}
		instances, err := discoverclient.DiscoverServices("Trip", "", true)
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
			"GenTrip" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("GenTrip")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &trip.GenTripRequest{
			PassengerReq: &trip.PublishOrderRequest{
				PassengerId:   1,
				StartTime:     0,
				Origin:        "三元里",
				Destination:   "成华大道",
				PassengerName: "张三",
			},
			DriverReq:    &trip.TakeOrderRequest{
				DriverId:   2,
				DriverName: "李四",
				Location:   "三里屯",
				Car:        "奔驰",
			},
		}
		for i := 0; i < 10; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			res, err := conn.GenTrip(ctx1, data)
			if err != nil {
				log.Printf("conn trip failed, err: %+v", err)
				return
			}
			fmt.Printf("%v %+v\n", res.Status, res.Trip)
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
			"GenTrip": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("GenTrip")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &trip.GenTripRequest{
			PassengerReq: &trip.PublishOrderRequest{
				PassengerId:   1,
				StartTime:     0,
				Origin:        "三元里",
				Destination:   "成华大道",
				PassengerName: "张三",
			},
			DriverReq:    &trip.TakeOrderRequest{
				DriverId:   2,
				DriverName: "李四",
				Location:   "三里屯",
				Car:        "奔驰",
			},
		}
		for i := 0; i < 10; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			r2, err := svc2.GenTrip(ctx1, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%v %+v\n", r2.Status, r2.Trip)
			//childSpan.Finish()
		}
	}
}
