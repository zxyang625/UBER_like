package endpoint

import (
	"context"
	"pkg/pb"
	service "trip/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GenTripRequest collects the request parameters for the GenTrip method.
type GenTripRequest struct {
	Req *pb.GenTripRequest `json:"req"`
}

// GenTripResponse collects the response parameters for the GenTrip method.
type GenTripResponse struct {
	Resp *pb.GenTripReply `json:"resp"`
	Err  error            `json:"err"`
}

// MakeGenTripEndpoint returns an endpoint that invokes GenTrip on the service.
func MakeGenTripEndpoint(s service.TripService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenTripRequest)
		resp, err := s.GenTrip(ctx, req.Req)
		return GenTripResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GenTripResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GenTrip implements Service. Primarily useful in a client.
func (e Endpoints) GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error) {
	request := GenTripRequest{Req: req}
	response, err := e.GenTripEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GenTripResponse).Resp, response.(GenTripResponse).Err
}
