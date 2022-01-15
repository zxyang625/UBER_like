package service

import (
	"context"
	grpc1 "driver/client/grpc"
	"encoding/json"
	"fmt"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	grpc2 "github.com/go-kit/kit/transport/grpc"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"log"
	"pkg/dao/models"
	"pkg/pb"
	"pkg/tracing"
)

// PassengerService describes the service.
type PassengerService interface {
	GetPassengerInfo(ctx context.Context, req *pb.GetPassengerInfoRequest) (resp *pb.GetPassengerInfoReply, err error)
	PublishOrder(ctx context.Context, req *pb.PublishOrderRequest) (resp *pb.PublishOrderReply, err error)
}
type basicPassengerService struct{}

func (b *basicPassengerService) GetPassengerInfo(ctx context.Context, req *pb.GetPassengerInfoRequest) (resp *pb.GetPassengerInfoReply, err error) {
	resp = &pb.GetPassengerInfoReply{}
	passenger, err := models.GetPassenger(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("GetPassengerInfo fail, err: %v", err)
	}
	data, err := json.Marshal(passenger)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	asset, _ := models.GetAsset(passenger.AccountNum)
	resp.Asset = asset
	return resp, nil
}

func (b *basicPassengerService) PublishOrder(ctx context.Context, req *pb.PublishOrderRequest) (resp *pb.PublishOrderReply, err error) {
	conn, err := grpc.Dial("127.0.0.1:8092", grpc.WithInsecure())
	if err != nil {
		log.Println("1", err)
		return
	}
	defer conn.Close()
	reporter := http.NewReporter(tracing.DefaultZipkinURL)
	defer reporter.Close()
	zkTracer, err := zipkingo.NewTracer(reporter)
	if err != nil {
		log.Printf("New GRPC TracingImpl failed, err: %v", err)
		return
	}
	zkClientTrace := kitzipkin.GRPCClientTrace(zkTracer)
	svc, err := grpc1.New(conn, map[string][]grpc2.ClientOption{
		"TakeOrder" : {zkClientTrace},
	})
	resp1, err := svc.TakeOrder(ctx, &pb.TakeOrderRequest{})
	if err != nil {
		return nil, err
	}
	log.Println("resp1", resp1.StartTime)
	resp = &pb.PublishOrderReply{
		Status:     true,
		DriverName: "老司机",
		Location:   "三元里",
		Car:        "北京现代",
		Path:       "直走就完事了",
	}
	return resp, nil
}

// NewBasicPassengerService returns a naive, stateless implementation of PassengerService.
func NewBasicPassengerService() PassengerService {
	return &basicPassengerService{}
}

// New returns a PassengerService with all of the expected middleware wired in.
func New(middleware []Middleware) PassengerService {
	var svc PassengerService = NewBasicPassengerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
