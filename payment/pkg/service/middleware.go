package service

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(PaymentService) PaymentService

type loggingMiddleware struct {
	logger log.Logger
	next   PaymentService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a PaymentService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PaymentService) PaymentService {
		return &loggingMiddleware{logger, next}
	}
}

func (l loggingMiddleware) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	defer func(start time.Time) {
		l.logger.Log("method", "Pay", "billNum", billNum, "accountNum", accountNum, "payPassword", payPassword, "msg", msg, "err", err)
	}(time.Now())
	return l.next.Pay(ctx, billNum, accountNum, payPassword)
}

