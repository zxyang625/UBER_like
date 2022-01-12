package service

import (
	"context"
	"passenger"
)

// PassengerService describes the service.
type PassengerService interface {
	GetPassengerInfo(ctx context.Context, req *passenger.GetPassengerInfoRequest) (resp *passenger.GetPassengerInfoReply, err error)
	PublishOrder(ctx context.Context, req *passenger.PublishOrderRequest) (resp *passenger.PublishOrderReply, err error)
}
type basicPassengerService struct{}

func (b *basicPassengerService) GetPassengerInfo(ctx context.Context, req *passenger.GetPassengerInfoRequest) (resp *passenger.GetPassengerInfoReply, err error) {
	resp = &passenger.GetPassengerInfoReply{
		UserId:     1111111111111,
		Name:       req.Username,
		Age:        12,
		AccountNum: 123456789,
		Asset:      5245.656,
	}
	return resp, nil
}
func (b *basicPassengerService) PublishOrder(ctx context.Context, req *passenger.PublishOrderRequest) (resp *passenger.PublishOrderReply, err error) {
	resp = &passenger.PublishOrderReply{
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
