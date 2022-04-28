package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"net/http"
	"pkg/dao/models"
	"pkg/dao/mq"
	Err "pkg/error"
	"pkg/pb"
	"time"
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
	traceID, _ := model.TraceIDFromHex(ctx.Value("Trace-ID").(string))
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("err1 %v", err)
	}
	asyncReq := mq.AsyncReq{
		Method:        http.MethodPost,
		OriginApp:     "passenger",
		OriginService: "publish-order",
		DestApp:       "trip",
		DestService:   "gen-trip",
		TraceID:       traceID,
		Priority:      ctx.Value("Length").(int),
		Header:        nil,
		Data:          data,
	}
	mqData, err := json.Marshal(asyncReq)
	if err != nil {
		return nil, fmt.Errorf("err2 %v", err)
	}
	err = PassengerMessageServer.Publish(ctx, PublishQueueName, ctx.Value("Length").(int), mqData)
	if err != nil {
		return nil, err
	}
	resp = &pb.PublishOrderReply{}
	c := make(chan struct{}, 1)
	d := amqp.Delivery{}
	go func() {
		d, err = PassengerMessageServer.ReceiveResp(ctx)
		c <- struct{}{}
	}()
	select {
	case <-c:
		if err != nil {
			return nil, fmt.Errorf("err3 %v", err)
		}
		err = json.Unmarshal(d.Body, resp)
		if err != nil {
			return nil, fmt.Errorf("err4 %v", err)
		}
		return
	case <-time.After(time.Second):
		return nil, Err.New(Err.RPCRequestTimeout, "PublishOrder timeout")
	}
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
