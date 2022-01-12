package service

import (
	"billing"
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(BillingService) BillingService

type loggingMiddleware struct {
	logger log.Logger
	next   BillingService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a BillingService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next BillingService) BillingService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GenBill(ctx context.Context, req *billing.GenBillRequest) (resp *billing.GenBillReply, err error) {
	defer func() {
		l.logger.Log("method", "GenBill", "req", req, "resp", resp, "err", err)
	}()
	return l.next.GenBill(ctx, req)
}
func (l loggingMiddleware) GetBillList(ctx context.Context, userId int64) (resp []*billing.BillMsg, err error) {
	defer func() {
		l.logger.Log("method", "GetBillList", "userId", userId, "resp", resp, "err", err)
	}()
	return l.next.GetBillList(ctx, userId)
}
