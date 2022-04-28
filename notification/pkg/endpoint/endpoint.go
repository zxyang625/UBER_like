package endpoint

import (
	"context"
	"pkg/pb"
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

// NoticeTrip implements OriginService. Primarily useful in a client.
func (e Endpoints) NoticeTrip(ctx context.Context, req *pb.NoticeTripRequest) (resp *pb.NoticeTripReply, err error) {
	request := NoticeTripRequest{Req: req}
	response, err := e.NoticeTripEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(NoticeTripResponse).Resp, response.(NoticeTripResponse).Err
}

// NoticeBillRequest collects the request parameters for the NoticeBill method.
type NoticeBillRequest struct {
	Req *pb.NoticeBillRequest `json:"req"`
}

// NoticeBillResponse collects the response parameters for the NoticeBill method.
type NoticeBillResponse struct {
	Resp *pb.NoticeBillReply `json:"resp"`
	Err  error               `json:"err"`
}

// MakeNoticeBillEndpoint returns an endpoint that invokes NoticeBill on the service.
func MakeNoticeBillEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(NoticeBillRequest)
		resp, err := s.NoticeBill(ctx, req.Req)
		return NoticeBillResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r NoticeBillResponse) Failed() error {
	return r.Err
}

// NoticeBill implements OriginService. Primarily useful in a client.
func (e Endpoints) NoticeBill(ctx context.Context, req *pb.NoticeBillRequest) (resp *pb.NoticeBillReply, err error) {
	request := NoticeBillRequest{Req: req}
	response, err := e.NoticeBillEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(NoticeBillResponse).Resp, response.(NoticeBillResponse).Err
}
