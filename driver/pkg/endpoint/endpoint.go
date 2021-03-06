package endpoint

import (
	"context"
	service "driver/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
	"pkg/pb"
)

// GetDriverInfoRequest collects the request parameters for the GetDriverInfo method.
type GetDriverInfoRequest struct {
	Req *pb.GetDriverInfoRequest `json:"req"`
}

// GetDriverInfoResponse collects the response parameters for the GetDriverInfo method.
type GetDriverInfoResponse struct {
	Resp *pb.GetDriverInfoReply `json:"resp"`
	Err  error                  `json:"err"`
}

// MakeGetDriverInfoEndpoint returns an endpoint that invokes GetDriverInfo on the service.
func MakeGetDriverInfoEndpoint(s service.DriverService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDriverInfoRequest)
		resp, err := s.GetDriverInfo(ctx, req.Req)
		return GetDriverInfoResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GetDriverInfoResponse) Failed() error {
	return r.Err
}

// TakeOrderRequest collects the request parameters for the TakeOrder method.
type TakeOrderRequest struct {
	Req *pb.TakeOrderRequest `json:"req"`
}

// TakeOrderResponse collects the response parameters for the TakeOrder method.
type TakeOrderResponse struct {
	Resp *pb.TakeOrderReply `json:"resp"`
	Err  error              `json:"err"`
}

// MakeTakeOrderEndpoint returns an endpoint that invokes TakeOrder on the service.
func MakeTakeOrderEndpoint(s service.DriverService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TakeOrderRequest)
		resp, err := s.TakeOrder(ctx, req.Req)
		return TakeOrderResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r TakeOrderResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetDriverInfo implements OriginService. Primarily useful in a client.
func (e Endpoints) GetDriverInfo(ctx context.Context, req *pb.GetDriverInfoRequest) (resp *pb.GetDriverInfoReply, err error) {
	request := GetDriverInfoRequest{Req: req}
	response, err := e.GetDriverInfoEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetDriverInfoResponse).Resp, response.(GetDriverInfoResponse).Err
}

// TakeOrder implements OriginService. Primarily useful in a client.
func (e Endpoints) TakeOrder(ctx context.Context, req *pb.TakeOrderRequest) (resp *pb.TakeOrderReply, err error) {
	request := TakeOrderRequest{Req: req}
	response, err := e.TakeOrderEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TakeOrderResponse).Resp, response.(TakeOrderResponse).Err
}
