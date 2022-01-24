package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"pkg/dao/models"
	"pkg/dao/mq"
	Err "pkg/error"
	"pkg/pb"
	"time"
)

// DriverService describes the service.
type DriverService interface {
	GetDriverInfo(ctx context.Context, req *pb.GetDriverInfoRequest) (resp *pb.GetDriverInfoReply, err error)
	TakeOrder(ctx context.Context, req *pb.TakeOrderRequest) (resp *pb.TakeOrderReply, err error)
}

type basicDriverService struct{}

func (b *basicDriverService) GetDriverInfo(ctx context.Context, req *pb.GetDriverInfoRequest) (resp *pb.GetDriverInfoReply, err error) {
	resp = &pb.GetDriverInfoReply{}
	driver, err := models.GetDriver(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("GetDriverInfo fail, err: %v", err)
	}
	data, err := json.Marshal(driver)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	asset, _ := models.GetAsset(driver.AccountNum)
	resp.Asset = asset
	return resp, nil
}

func (b *basicDriverService) TakeOrder(ctx context.Context, req *pb.TakeOrderRequest) (resp *pb.TakeOrderReply, err error) {
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	span := zipkin.SpanOrNoopFromContext(ctx)
	mqModel := mq.MQModel{
		Data: data,
		SpanModel: model.SpanModel{
			SpanContext:    model.SpanContext{
				TraceID:  span.Context().TraceID,
				ID:       span.Context().ID,
				ParentID: span.Context().ParentID,
			},
		},
	}
	mqData, err := json.Marshal(mqModel)
	if err != nil {
		return nil, err
	}
	err = DriverMessageServer.Publish(ctx, PublishQueueName, mqData)
	if err != nil {
		return nil, err
	}
	resp = &pb.TakeOrderReply{}
	c := make(chan struct{}, 1)
	d := amqp.Delivery{}
	go func() {
		d, err = DriverMessageServer.ReceiveResp(ctx)
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
		return nil, Err.New(Err.RPCRequestTimeout, "TakeOrder timeout")
	}
}

// NewBasicDriverService returns a naive, stateless implementation of DriverService.
func NewBasicDriverService() DriverService {
	return &basicDriverService{}
}

// New returns a DriverService with all of the expected middleware wired in.
func New(middleware []Middleware) DriverService {
	var svc DriverService = NewBasicDriverService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
