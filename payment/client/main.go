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
	grpc1 "payment/client/grpc"
	"pkg/discover"
	"pkg/zipkin"
)

func main() {
	discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
	instances, err := discoverclient.DiscoverServices("Payment", "", true)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(discover.GetInstance(instances))

	///////////////////////////////////
	reporter := http.NewReporter(zipkin.DefaultZipkinURL)	//zipkin
	zkTracer, err := zipkingo.NewTracer(reporter)
	if err != nil {
		log.Printf("NewTracer failed, err: %v", err)
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
		"Pay" : {zkClientTrace},
	})
	if err != nil {
		log.Println("2", err)
		return
	}
	parentSpan := zkTracer.StartSpan("Pay")
	defer parentSpan.Flush()
	ctx := zipkingo.NewContext(context.Background(), parentSpan)
	for i := 0; i < 100; i++ {
		r2, _ := svc2.Pay(ctx, 3124,1514, "34rfey345re")
		fmt.Println(r2)
	}
}
