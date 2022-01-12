package grpc

import (
	"context"
	"driver"
	endpoint "driver/pkg/endpoint"
	pb "driver/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetDriverInfoHandler creates the handler logic
func makeGetDriverInfoHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetDriverInfoEndpoint, decodeGetDriverInfoRequest, encodeGetDriverInfoResponse, options...)
}

// decodeGetDriverInfoResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetDriverInfo request.
// TODO implement the decoder
func decodeGetDriverInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req :=r.(*pb.GetDriverInfoRequest)
	return endpoint.GetDriverInfoRequest{
		Req: &driver.DriverInfoRequest{
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	}, nil
}

// encodeGetDriverInfoResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetDriverInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetDriverInfoResponse)
	return &pb.GetDriverInfoReply{
		UserId:               resp.Resp.UserId,
		Name:                 resp.Resp.Name,
		Age:                  resp.Resp.Age,
		AccountNum:           resp.Resp.AccountNum,
	}, resp.Err
}
func (g *grpcServer) GetDriverInfo(ctx context1.Context, req *pb.GetDriverInfoRequest) (*pb.GetDriverInfoReply, error) {
	_, rep, err := g.getDriverInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetDriverInfoReply), nil
}

// makeTakeOrderHandler creates the handler logic
func makeTakeOrderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.TakeOrderEndpoint, decodeTakeOrderRequest, encodeTakeOrderResponse, options...)
}

// decodeTakeOrderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain TakeOrder request.
// TODO implement the decoder
func decodeTakeOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.TakeOrderRequest)
	return endpoint.TakeOrderRequest{
		Req: &driver.TakeOrderRequest{
			DriverId:   req.GetDriverId(),
			DriverName: req.GetDriverName(),
			Location:   req.GetLocation(),
			Car:        req.GetCar(),
		},
	}, nil
}

// encodeTakeOrderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeTakeOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.TakeOrderResponse)
	return &pb.TakeOrderReply{
		PassengerName:        resp.Resp.PassengerName,
		StartTime:            resp.Resp.StartTime,
		Origin:               resp.Resp.Origin,
		Destination:          resp.Resp.Destination,
		Path:                 resp.Resp.Path,
	}, resp.Err
}
func (g *grpcServer) TakeOrder(ctx context1.Context, req *pb.TakeOrderRequest) (*pb.TakeOrderReply, error) {
	_, rep, err := g.takeOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TakeOrderReply), nil
}
