package grpc

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "passenger/pkg/endpoint"
	service "passenger/pkg/service"
	"pkg/pb"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.PassengerService, error) {
	var getPassengerInfoEndpoint endpoint.Endpoint
	{
		getPassengerInfoEndpoint = grpc1.NewClient(conn, "pb.Passenger", "GetPassengerInfo", encodeGetPassengerInfoRequest, decodeGetPassengerInfoResponse, pb.GetPassengerInfoReply{}, options["GetPassengerInfo"]...).Endpoint()
	}

	var publishOrderEndpoint endpoint.Endpoint
	{
		publishOrderEndpoint = grpc1.NewClient(conn, "pb.Passenger", "PublishOrder", encodePublishOrderRequest, decodePublishOrderResponse, pb.PublishOrderReply{}, options["PublishOrder"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		GetPassengerInfoEndpoint: getPassengerInfoEndpoint,
		PublishOrderEndpoint:     publishOrderEndpoint,
	}, nil
}

// encodeGetPassengerInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetPassengerInfo request to a gRPC request.
func encodeGetPassengerInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.GetPassengerInfoRequest)
	return &pb.GetPassengerInfoRequest{
		Username:             req.Req.Username,
		Password:             req.Req.Password,
	}, nil
}

// decodeGetPassengerInfoResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetPassengerInfoResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.GetPassengerInfoReply)
	return endpoint1.GetPassengerInfoResponse{
		Resp: &pb.GetPassengerInfoReply{
			UserId:     resp.GetUserId(),
			Name:       resp.GetName(),
			Age:        resp.GetAge(),
			AccountNum: resp.GetAccountNum(),
			Asset:      resp.GetAsset(),
		},
		Err:  nil,
	}, nil
}

// encodePublishOrderRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain PublishOrder request to a gRPC request.
func encodePublishOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.PublishOrderRequest)
	return &pb.PublishOrderRequest{
		PassengerId:          req.Req.PassengerId,
		StartTime:            req.Req.StartTime,
		Origin:               req.Req.Origin,
		Destination:          req.Req.Destination,
		PassengerName:        req.Req.PassengerName,
	}, nil
}

// decodePublishOrderResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodePublishOrderResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.PublishOrderReply)
	return endpoint1.PublishOrderResponse{
		Resp: &pb.PublishOrderReply{
			Status:     resp.GetStatus(),
			DriverName: resp.GetDriverName(),
			Location:   resp.GetLocation(),
			Car:        resp.GetCar(),
			Path:       resp.GetPath(),
		},
		Err: nil,
	}, nil
}
