package grpc

import (
	"context"
	"driver"
	endpoint1 "driver/pkg/endpoint"
	pb "driver/pkg/grpc/pb"
	service "driver/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.DriverService, error) {
	var getDriverInfoEndpoint endpoint.Endpoint
	{
		getDriverInfoEndpoint = grpc1.NewClient(conn, "pb.Driver", "GetDriverInfo", encodeGetDriverInfoRequest, decodeGetDriverInfoResponse, pb.GetDriverInfoReply{}, options["GetDriverInfo"]...).Endpoint()
	}

	var takeOrderEndpoint endpoint.Endpoint
	{
		takeOrderEndpoint = grpc1.NewClient(conn, "pb.Driver", "TakeOrder", encodeTakeOrderRequest, decodeTakeOrderResponse, pb.TakeOrderReply{}, options["TakeOrder"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		GetDriverInfoEndpoint: getDriverInfoEndpoint,
		TakeOrderEndpoint:     takeOrderEndpoint,
	}, nil
}

// encodeGetDriverInfoRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetDriverInfo request to a gRPC request.
func encodeGetDriverInfoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.GetDriverInfoRequest)
	return &pb.GetDriverInfoRequest{
		Username: req.Req.Username,
		Password: req.Req.Password,
	}, nil
}

// decodeGetDriverInfoResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetDriverInfoResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.GetDriverInfoReply)
	return endpoint1.GetDriverInfoResponse{
		Resp: &driver.DriverInfoReply{
			UserId:     resp.UserId,
			Name:       resp.Name,
			Age:        resp.Age,
			AccountNum: resp.AccountNum,
			Asset:      resp.Asset,
		},
		Err: nil,
	}, nil
}

// encodeTakeOrderRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain TakeOrder request to a gRPC request.
func encodeTakeOrderRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.TakeOrderRequest)
	return &pb.TakeOrderRequest{
		DriverId:             req.Req.DriverId,
		DriverName:           req.Req.DriverName,
		Location:             req.Req.Location,
		Car:                  req.Req.Car,
	}, nil
}

// decodeTakeOrderResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeTakeOrderResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.TakeOrderReply)
	return endpoint1.TakeOrderResponse{
		Resp: &driver.TakeOrderReply{
			PassengerName: resp.PassengerName,
			StartTime:     resp.StartTime,
			Origin:        resp.Origin,
			Destination:   resp.Destination,
			Path:          resp.Path,
		},
		Err:  nil,
	}, nil
}
