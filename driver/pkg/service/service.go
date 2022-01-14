package service

import (
	"context"
	"encoding/json"
	"fmt"
	"pkg/dao/models"
	"pkg/pb"
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
	return &pb.TakeOrderReply{
		PassengerName: "李四",
		StartTime:     11111111,
		Origin:        "三里屯",
		Destination:   "",
		Path:          "",
	}, nil
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
