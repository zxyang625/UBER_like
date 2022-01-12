package service

import (
	"context"
	"driver"
	"time"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(DriverService) DriverService

type loggingMiddleware struct {
	logger log.Logger
	next   DriverService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PaymentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DriverService) DriverService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) GetDriverInfo(ctx context.Context, req *driver.DriverInfoRequest) (resp *driver.DriverInfoReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "GetDriverInfo", "err", err)
	}(time.Now())
	return l.next.GetDriverInfo(ctx, req)
}

func (l loggingMiddleware) TakeOrder(ctx context.Context, req *driver.TakeOrderRequest) (resp *driver.TakeOrderReply, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "TakeOrder", "err", err)
	}(time.Now())
	return l.next.TakeOrder(ctx, req)
}
