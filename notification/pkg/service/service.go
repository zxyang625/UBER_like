package service

import (
	"context"
	"pkg/pb"
)

// NotificationService describes the service.
type NotificationService interface {
	NoticeTrip(ctx context.Context, req *pb.NoticeTripRequest) (resp *pb.NoticeTripReply, err error)
	NoticeBill(ctx context.Context, req *pb.NoticeBillRequest) (resp *pb.NoticeBillReply, err error)
}

type basicNotificationService struct{}

func (b *basicNotificationService) NoticeTrip(ctx context.Context, req *pb.NoticeTripRequest) (resp *pb.NoticeTripReply, err error) {
	return &pb.NoticeTripReply{
		Status: true,
		Msg:    "success" + req.GetTripMsg().GetPassengerName() + req.GetTripMsg().GetCar(),
	}, nil
}

func (b *basicNotificationService) NoticeBill(ctx context.Context, req *pb.NoticeBillRequest) (resp *pb.NoticeBillReply, err error) {
	return &pb.NoticeBillReply{
		Status: true,
		Msg:    "success" + req.GetBillMsg().GetPassengerName() + req.GetBillMsg().GetDriverName(),
	}, nil
}

// NewBasicNotificationService returns a naive, stateless implementation of NotificationService.
func NewBasicNotificationService() NotificationService {
	return &basicNotificationService{}
}

// New returns a NotificationService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificationService {
	var svc NotificationService = NewBasicNotificationService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
