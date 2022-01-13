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
	grpc1 "notification/client/grpc"
	http1 "notification/client/http"
	"notification/pkg/grpc/pb"
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
		instances, err := discoverclient.DiscoverServices("Notification", "", true)
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
			"NoticeBill" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Notification")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &pb.NoticeBillRequest{
			BillMsg:              &pb.BillMsg{PassengerName: "王五", DriverName: "赵四"},
		}
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			res, err := conn.NoticeBill(ctx1, data)
			if err != nil {
				log.Printf("conn NoticeBill failed, err: %v", err)
				return
			}
			fmt.Println(res.GetMsg(), res.Status)
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
			"NoticeBill": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("Notification")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &pb.NoticeBillRequest{
				BillMsg:    &pb.BillMsg{
					PassengerName:        "张三",
				},
		}
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			r2, err := svc2.NoticeBill(ctx1, data)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(r2.Msg, r2.Status)
			//childSpan.Finish()
		}
	}
}
