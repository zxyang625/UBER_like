package main

import (
	"billing"
	grpc1 "billing/client/grpc"
	http1 "billing/client/http"
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
)

func main() {
	{	//discover
		discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
		if err != nil {
			log.Printf("1, err: %v\n", err)
			return
		}
		instances, err := discoverclient.DiscoverServices("Billing", "", true)
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
			"GenBill" : {zkClientTrace},
		})
		if err != nil {
			log.Printf("new http conn failed, err: %v", err)
			return
		}
		parentSpan := zkTracer.StartSpan("GenBill")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := &billing.GenBillRequest{
			TripMsg: &billing.TripMsg{
			TripNum:       123456789,
			PassengerId:   123,
			DriverId:      5234,
			PassengerName: "req.PassengerReq.PassengerName",
			DriverName:    "req.DriverReq.DriverName",
			StartTime:     12,
			EndTime:       18,
			Origin:        "req.PassengerReq.Origin",
			Destination:   "req.PassengerReq.Destination",
			Car:           "req.DriverReq.Car",
			Path:          "直走一公里后右转",
		}}
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			res, err := conn.GenBill(ctx1, data)
			if err != nil {
				log.Printf("conn Billing failed, err: %v", err)
				return
			}
			fmt.Printf("%v %+v\n", res.Status, res.BillMsg)
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
			"GetBillList": {zkClientTrace},
		})
		if err != nil {
			log.Println("2", err)
			return
		}
		parentSpan := zkTracer.StartSpan("GetBillList")
		defer parentSpan.Finish()
		ctx := zipkingo.NewContext(context.Background(), parentSpan)
		data := 1232154156
		for i := 0; i < 100; i++ {
			//childSpan := zkTracer.StartSpan("childSpan", zipkingo.Parent(parentSpan.Context()))
			ctx1 := zipkingo.NewContext(ctx, parentSpan)
			r2, err := svc2.GetBillList(ctx1, int64(data))
			if err != nil {
				log.Println(err)
				return
			}
			for i := range r2 {
				fmt.Printf("%+v\n", r2[i])
			}
			//childSpan.Finish()
		}
	}
}

