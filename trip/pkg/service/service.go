package service

import (
	"context"
	"trip"
)

// TripService describes the service.
type TripService interface {
	GenTrip(ctx context.Context, req *trip.GenTripRequest) (resp *trip.GenTripReply, err error)
}

type basicTripService struct{}

func (b *basicTripService) GenTrip(ctx context.Context, req *trip.GenTripRequest) (resp *trip.GenTripReply, err error) {
	return &trip.GenTripReply{
		Status: true,
		Trip: &trip.TripMsg{
			TripNum:       123456789,
			PassengerId:   req.PassengerReq.PassengerId,
			DriverId:      req.DriverReq.DriverId,
			PassengerName: req.PassengerReq.PassengerName,
			DriverName:    req.DriverReq.DriverName,
			StartTime:     12,
			EndTime:       18,
			Origin:        req.PassengerReq.Origin,
			Destination:   req.PassengerReq.Destination,
			Car:           req.DriverReq.Car,
			Path:          "直走一公里后右转",
		},
	}, nil
}

// NewBasicTripService returns a naive, stateless implementation of TripService.
func NewBasicTripService() TripService {
	return &basicTripService{}
}

// New returns a TripService with all of the expected middleware wired in.
func New(middleware []Middleware) TripService {
	var svc TripService = NewBasicTripService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
