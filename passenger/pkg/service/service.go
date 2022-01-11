package service

import "context"

type GetPassengerInfoRequest struct {
	Username string
	Password string
}

type GetPassengerInfoReply struct {
	userId     int64
	Name       string
	Age        int32
	AccountNum int64
	Asset      float64
}

type PublishRequest struct {
	PassengerId   int64
	StartTime     int64
	Origin        string
	Destination   string
	PassengerName string
}

type PublishReply struct {
	status     bool
	DriverName string
	Location   string
	Car        string
	Path       string
}

// PassengerService describes the service.
type PassengerService interface {
	GetPassengerInfo(ctx context.Context, req *GetPassengerInfoRequest) (resp *GetPassengerInfoReply, err error)
	PublishOrder(ctx context.Context, req *PublishRequest) (resp *PublishReply, err error)
}
