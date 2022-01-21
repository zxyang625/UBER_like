package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"pkg/pb"
	"time"
)

// Middleware describes a service middleware.
type Middleware func(PassengerService) PassengerService

type loggingMiddleware struct {
	logger log.Logger
	next   PassengerService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PaymentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PassengerService) PassengerService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) GetPassengerInfo(ctx context.Context, req *pb.GetPassengerInfoRequest) (resp *pb.GetPassengerInfoReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "GetPassengerInfo", "err", err)
	}(time.Now())
	return l.next.GetPassengerInfo(ctx, req)
}

func (l loggingMiddleware) PublishOrder(ctx context.Context, req *pb.PublishOrderRequest) (resp *pb.PublishOrderReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "PublishOrder", "err", err)
	}(time.Now())
	return l.next.PublishOrder(ctx, req)
}
