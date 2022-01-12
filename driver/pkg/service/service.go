package service

import (
	"context"
	"driver"
	"errors"
)

// DriverService describes the service.
type DriverService interface {
	GetDriverInfo(ctx context.Context, req *driver.DriverInfoRequest) (resp *driver.DriverInfoReply, err error)
	TakeOrder(ctx context.Context, req *driver.TakeOrderRequest) (resp *driver.TakeOrderReply, err error)
}

type basicDriverService struct{}

func (b *basicDriverService) GetDriverInfo(ctx context.Context, req *driver.DriverInfoRequest) (resp *driver.DriverInfoReply, err error) {
	return &driver.DriverInfoReply{
		UserId:     123456,
		Name:       "张三",
		Age:        0,
		AccountNum: 0,
		Asset:      0,
	}, errors.New("http error")
}
func (b *basicDriverService) TakeOrder(ctx context.Context, req *driver.TakeOrderRequest) (resp *driver.TakeOrderReply, err error) {
	return &driver.TakeOrderReply{
		PassengerName: "李四",
		StartTime:     11111111,
		Origin:        "三里屯",
		Destination:   "",
		Path:          "",
	}, errors.New("grpc error")
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
