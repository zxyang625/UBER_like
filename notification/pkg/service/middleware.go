package service

import (
	"context"
	"pkg/pb"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificationService) NotificationService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificationService) NotificationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) NoticeTrip(ctx context.Context, req *pb.NoticeTripRequest) (resp *pb.NoticeTripReply, err error) {
	defer func() {
		l.logger.Log("method", "NoticeTrip", "TripNum", req.TripMsg.GetTripNum(), "err", err)
	}()
	return l.next.NoticeTrip(ctx, req)
}

func (l loggingMiddleware) NoticeBill(ctx context.Context, req *pb.NoticeBillRequest) (resp *pb.NoticeBillReply, err error) {
	defer func() {
		l.logger.Log("method", "NoticeBill", "BillNum", req.BillMsg.GetBillNum(), "err", err)
	}()
	return l.next.NoticeBill(ctx, req)
}
