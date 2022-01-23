package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"pkg/dao/models"
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
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = PassengerMessageServer.Publish(ctx, PublishQueueName, data)
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
			return nil, err
		}
		err = proto.Unmarshal(d.Body, resp)
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
