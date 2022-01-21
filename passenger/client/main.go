package main

import (
	"context"
	"fmt"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
	grpc1 "passenger/client/grpc"
	"pkg/pb"
	"pkg/tracing"
)

func main() {
	//{	//discover
	//	discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
	//	if err != nil {
	//		log.Printf("1, err: %v\n", err)
	//		return
	//	}
	//	instances, err := discoverclient.DiscoverServices("Passenger", "", true)
	//	if err != nil {
	//		log.Printf("2, err: %v\n", err)
	//		return
	//	}
	//	fmt.Println(discover.GetInstance(instances))
	//}
	//
	//{	//http
	//	reporter := http.NewReporter(tracing.DefaultZipkinURL) //zipkin
	//	defer reporter.Close()
	//	zkTracer, err := zipkingo.NewTracer(reporter)
	//	if err != nil {
	//		log.Printf("New HTTP TracingImpl failed, err: %v", err)
	//		return
	//	}
	//	zkClientTrace := kitzipkin.HTTPClientTrace(zkTracer)
	//	conn, err := http1.New("127.0.0.1:8081", map[string][]kithttp.ClientOption{
	//		"GetPassengerInfo" : {zkClientTrace},
	//	})
	//	if err != nil {
	//		log.Printf("new http conn failed, err: %v", err)
	//		return
	//	}
	//	parentSpan := zkTracer.StartSpan("Passenger")
	//	defer parentSpan.Finish()
	//	ctx := zipkingo.NewContext(context.Background(), parentSpan)
	//	data := &pb.GetPassengerInfoRequest{
	//		Username: "张三",
	//		Password: "123456",
	//	}
	//	//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
	//	ctx1 := zipkingo.NewContext(ctx, parentSpan)
	//	res, err := conn.GetPassengerInfo(ctx1, data)
	//	if err != nil {
	//		log.Printf("conn Passenger failed, err: %v", err)
	//		return
	//	}
	//	fmt.Println(res.Asset, res.Age, res.Name)
	//	//childSpan.Finish()
	//}

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
			"GetPassengerInfo": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		data := &pb.PublishOrderRequest{
			PassengerId:   3124,
			StartTime:     52,
			Origin:        "北京三元里",
			Destination:   "天安门广场",
			PassengerName: "张三",
		}
		//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
		//ctx1 := zipkingo.NewContext(ctx, parentSpan)
		parentSpan := zkTracer.StartSpan("Passenger_request")
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		defer parentSpan.Finish()
		//go func() {
		//	for i := 0; i < 10000; i++ {
		//		r2, err := svc2.PublishOrder(ctx, data)
		//		if err != nil {
		//			log.Println(err)
		//			//return
		//		}
		//		fmt.Println(r2.GetStatus(), r2.GetMsg())
		//		time.Sleep(time.Second / 100)
		//	}
		//}()
		//go func() {
		//	for i := 0; i < 10000; i++ {
		//		r2, err := svc2.PublishOrder(ctx, data)
		//		if err != nil {
		//			log.Println(err)
		//			//return
		//		}
		//		fmt.Println(r2.GetStatus(), r2.GetMsg())
		//		time.Sleep(time.Second / 100)
		//	}
		//}()
		for i := 0; i < 100000; i++ {
			r2, err := svc2.PublishOrder(ctx, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(r2.GetStatus(), r2.GetMsg())
			//time.Sleep(time.Second / 1000)
		}
		//childSpan.Finish()
	}
}
