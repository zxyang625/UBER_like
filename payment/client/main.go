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
	"pkg/tracing"
)

func main() {
	//{	//discover
	//	discoverclient, err := discover.NewDiscoverClient("127.0.0.1", 8500, true)
	//	if err != nil {
	//		log.Printf("1, err: %v\n", err)
	//		return
	//	}
	//	instances, err := discoverclient.DiscoverServices("Payment", "", true)
	//	if err != nil {
	//		log.Printf("2, err: %v\n", err)
	//		return
	//	}
	//	fmt.Println(discover.GetInstance(instances))
	//}
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
	//		"Pay" : {zkClientTrace},
	//	})
	//	if err != nil {
	//		log.Printf("new http conn failed, err: %v", err)
	//		return
	//	}
	//	parentSpan := zkTracer.StartSpan("Pay")
	//	defer parentSpan.Finish()
	//	ctx := zipkingo.NewContext(context.Background(), parentSpan)
	//	for i := 0; i < 100; i++ {
	//		res, err := conn.Pay(ctx, 9999, 9999, "253rfe64tgrw")
	//		if err != nil {
	//			log.Printf("conn Pay failed, err: %v", err)
	//			return
	//		}
	//		fmt.Println(res)
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
		r2, err := svc2.Pay(ctx, 1, 1, "qwer")
		if err != nil {
			fmt.Println("3", err)
			return
		}
		fmt.Println(r2)

	}
	{	//mysql
		//r1, err := models.GetAccount(1, "qwer")
		//fmt.Println(r1, err)
		////models.AddAccount(&models.Account{
		////	AccountNum:  2,
		////	PayPassword: "asdf",
		////	Asset:       354.6,
		////})
		////models.DelAccount(2)
		//models.UpdateAccount(1, &models.Account{Asset: 321})
		//
		//r2, _ := models.GetBill(1)
		//fmt.Println(r2)
		//models.AddBill(&models.Bill{
		//	BillNum:       2,
		//	Price:         65.7,
		//	StartTime:     123,
		//	EndTime:       556,
		//	Origin:        "三元里",
		//	Destination:   "成华大道",
		//	PassengerName: "张三",
		//	DriverName:    "李四",
		//	Payed:         false,
		//})
		//models.DelBill(2)
		//models.UpdateBill(1, &models.Bill{
		//	Origin: "二仙桥",
		//})
	}
}
