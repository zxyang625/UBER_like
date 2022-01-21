package service

import (
	"context"
	"pkg/pb"

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

func (l loggingMiddleware) GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error) {
	defer func() {
		l.logger.Log("method", "GenBill", "BillNum", resp.BillMsg.GetBillNum(), "err", err)
	}()
	return l.next.GenBill(ctx, req)
}
func (l loggingMiddleware) GetBillList(ctx context.Context, userId int64) (resp []*pb.BillMsg, err error) {
	defer func() {
		l.logger.Log("method", "GetBillList", "userId", userId, "err", err)
	}()
	return l.next.GetBillList(ctx, userId)
}

func (l loggingMiddleware) GetBill(ctx context.Context, billNum int64) (resp *pb.BillMsg, err error) {
	defer func() {
		l.logger.Log("method", "GetBill", "billNum", billNum, "resp", resp, "err", err)
	}()
	return l.next.GetBill(ctx, billNum)
}
