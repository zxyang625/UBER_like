package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	"pkg/pb"
)

// Middleware describes a service middleware.
type Middleware func(TripService) TripService

type loggingMiddleware struct {
	logger log.Logger
	next   TripService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TripService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TripService) TripService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error) {
	defer func() {
		l.logger.Log("method", "GenTrip", "TripNum", resp.GetTripMsg().GetTripNum(), "err", err)
	}()
	return l.next.GenTrip(ctx, req)
}
