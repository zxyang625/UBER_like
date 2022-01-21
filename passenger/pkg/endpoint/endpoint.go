package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "passenger/pkg/service"
	"pkg/pb"
)

// GetPassengerInfoRequest collects the request parameters for the GetPassengerInfo method.
type GetPassengerInfoRequest struct {
	Req *pb.GetPassengerInfoRequest `json:"req"`
}

// GetPassengerInfoResponse collects the response parameters for the GetPassengerInfo method.
type GetPassengerInfoResponse struct {
	Resp *pb.GetPassengerInfoReply `json:"resp"`
	Err  error                     `json:"err"`
}

// MakeGetPassengerInfoEndpoint returns an endpoint that invokes GetPassengerInfo on the service.
func MakeGetPassengerInfoEndpoint(s service.PassengerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPassengerInfoRequest)
		resp, err := s.GetPassengerInfo(ctx, req.Req)
		return GetPassengerInfoResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GetPassengerInfoResponse) Failed() error {
	return r.Err
}

// PublishOrderRequest collects the request parameters for the PublishOrder method.
type PublishOrderRequest struct {
	Req *pb.PublishOrderRequest `json:"req"`
}

// PublishOrderResponse collects the response parameters for the PublishOrder method.
type PublishOrderResponse struct {
	Resp *pb.PublishOrderReply `json:"resp"`
	Err  error                 `json:"err"`
}

// MakePublishOrderEndpoint returns an endpoint that invokes PublishOrder on the service.
func MakePublishOrderEndpoint(s service.PassengerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PublishOrderRequest)
		resp, err := s.PublishOrder(ctx, req.Req)
		return PublishOrderResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r PublishOrderResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetPassengerInfo implements Service. Primarily useful in a client.
func (e Endpoints) GetPassengerInfo(ctx context.Context, req *pb.GetPassengerInfoRequest) (resp *pb.GetPassengerInfoReply, err error) {
	request := GetPassengerInfoRequest{Req: req}
	response, err := e.GetPassengerInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetPassengerInfoResponse).Resp, response.(GetPassengerInfoResponse).Err
}

// PublishOrder implements Service. Primarily useful in a client.
func (e Endpoints) PublishOrder(ctx context.Context, req *pb.PublishOrderRequest) (resp *pb.PublishOrderReply, err error) {
	request := PublishOrderRequest{Req: req}
	response, err := e.PublishOrderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PublishOrderResponse).Resp, response.(PublishOrderResponse).Err
}
