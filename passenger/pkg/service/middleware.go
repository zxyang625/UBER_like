package service

import (
	"context"
	"passenger"
	"time"

	"github.com/go-kit/kit/log"
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

func (l loggingMiddleware) GetPassengerInfo(ctx context.Context, req *passenger.GetPassengerInfoRequest) (resp *passenger.GetPassengerInfoReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "GetPassengerInfo", "err", err)
	}(time.Now())
	return l.next.GetPassengerInfo(ctx, req)
}

func (l loggingMiddleware) PublishOrder(ctx context.Context, req *passenger.PublishOrderRequest) (resp *passenger.PublishOrderReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "GetPassengerInfo", "err", err)
	}(time.Now())
	return l.next.PublishOrder(ctx, req)
}
