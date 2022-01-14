package grpc

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	"pkg/pb"
	endpoint1 "trip/pkg/endpoint"
	service "trip/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.TripService, error) {
	var genTripEndpoint endpoint.Endpoint
	{
		genTripEndpoint = grpc1.NewClient(conn, "pb.Trip", "GenTrip", encodeGenTripRequest, decodeGenTripResponse, pb.GenTripReply{}, options["GenTrip"]...).Endpoint()
	}

	return endpoint1.Endpoints{GenTripEndpoint: genTripEndpoint}, nil
}

// encodeGenTripRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenTrip request to a gRPC request.
func encodeGenTripRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.GenTripRequest)
	return &pb.GenTripRequest{
		PassengerReq:         &pb.PublishOrderRequest{
			PassengerId:          req.Req.PassengerReq.PassengerId,
			StartTime:            req.Req.PassengerReq.StartTime,
			Origin:               req.Req.PassengerReq.Origin,
			Destination:          req.Req.PassengerReq.Destination,
			PassengerName:        req.Req.PassengerReq.PassengerName,
		},
		DriverReq:            &pb.TakeOrderRequest{
			DriverId:             req.Req.DriverReq.DriverId,
			DriverName:           req.Req.DriverReq.DriverName,
			Location:             req.Req.DriverReq.Location,
			Car:                  req.Req.DriverReq.Car,
		},
	}, nil
}

// decodeGenTripResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenTripResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.GenTripReply)
	return endpoint1.GenTripResponse{
		Resp: &pb.GenTripReply{
			Status: resp.GetStatus(),
			TripMsg:   &pb.TripMsg{
				TripNum:       resp.GetTripMsg().GetTripNum(),
				PassengerId:   resp.GetTripMsg().GetPassengerId(),
				DriverId:      resp.GetTripMsg().GetDriverId(),
				PassengerName: resp.GetTripMsg().GetPassengerName(),
				DriverName:    resp.GetTripMsg().GetDriverName(),
				StartTime:     resp.GetTripMsg().GetStartTime(),
				EndTime:       resp.GetTripMsg().GetEndTime(),
				Origin:        resp.GetTripMsg().GetOrigin(),
				Destination:   resp.GetTripMsg().GetDestination(),
				Car:           resp.GetTripMsg().GetCar(),
				Path:          resp.GetTripMsg().GetPath(),
			},
		},
		Err:  nil,
	}, nil
}
