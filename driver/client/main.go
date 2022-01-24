package main

import (
	"context"
	grpc1 "driver/client/grpc"
	"fmt"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
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
	//	instances, err := discoverclient.DiscoverServices("Driver", "", true)
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
	//		"GetDriverInfo" : {zkClientTrace},
	//	})
	//	if err != nil {
	//		log.Printf("new http conn failed, err: %v", err)
	//		return
	//	}
	//	parentSpan := zkTracer.StartSpan("Driver")
	//	defer parentSpan.Finish()
	//	ctx := zipkingo.NewContext(context.Background(), parentSpan)
	//	data := &pb.GetDriverInfoRequest{
	//		Username: "aaaaaaaaaaaa",
	//		Password: "bbbbbbbbbb",
	//	}
	//	for i := 0; i < 100; i++ {
	//		childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
	//		res, err := conn.GetDriverInfo(ctx, data)
	//		if err != nil {
	//			log.Printf("conn Driver failed, err: %v", err)
	//			return
	//		}
	//		fmt.Println(res)
	//		childSpan.Finish()
	//	}
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
		conn, err := grpc.Dial("127.0.0.1:8092", grpc.WithInsecure())
		if err != nil {
			log.Println("1", err)
			return
		}
		defer conn.Close()
		svc2, err := grpc1.New(conn, map[string][]grpc2.ClientOption{
			"TakeOrder": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		data := &pb.TakeOrderRequest{
			DriverId:   1234556,
			DriverName: "王老五",
			Location:   "街道口",
			Car:        "奔驰",
		}
		parentSpan := zkTracer.StartSpan("Driver")
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		defer parentSpan.Finish()
		//go func() {
		//	for i := 0; i < 10000; i++ {
		//		r2, err := svc2.TakeOrder(ctx, data)
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
		//		r2, err := svc2.TakeOrder(ctx, data)
		//		if err != nil {
		//			log.Println(err)
		//			//return
		//		}
		//		fmt.Println(r2.GetStatus(), r2.GetMsg())
		//		time.Sleep(time.Second / 100)
		//	}
		//}()
		//for i := 0; i < 100000; i++ {
			r2, err := svc2.TakeOrder(ctx, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(r2.GetStatus(), r2.GetMsg())
			//time.Sleep(time.Second / 1000)
		//}
	}
}
