package endpoint

import (
	"context"
	"notification/pkg/grpc/pb"
	service "notification/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// NoticeTripRequest collects the request parameters for the NoticeTrip method.
type NoticeTripRequest struct {
	Req *pb.NoticeTripRequest `json:"req"`
}

// NoticeTripResponse collects the response parameters for the NoticeTrip method.
type NoticeTripResponse struct {
	Resp *pb.NoticeTripReply `json:"resp"`
	Err  error               `json:"err"`
}

// MakeNoticeTripEndpoint returns an endpoint that invokes NoticeTrip on the service.
func MakeNoticeTripEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NoticeTripRequest)
		resp, err := s.NoticeTrip(ctx, req.Req)
		return NoticeTripResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r NoticeTripResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// NoticeTrip implements Service. Primarily useful in a client.
func (e Endpoints) NoticeTrip(ctx context.Context, req *pb.NoticeTripRequest) (resp *pb.NoticeTripReply, err error) {
	request := NoticeTripRequest{Req: req}
	response, err := e.NoticeTripEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(NoticeTripResponse).Resp, response.(NoticeTripResponse).Err
}
