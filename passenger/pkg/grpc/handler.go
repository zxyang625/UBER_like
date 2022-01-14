package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "passenger/pkg/endpoint"
	pb "pkg/pb"
)

// makeGetPassengerInfoHandler creates the handler logic
func makeGetPassengerInfoHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetPassengerInfoEndpoint, decodeGetPassengerInfoRequest, encodeGetPassengerInfoResponse, options...)
}

// decodeGetPassengerInfoResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetPassengerInfo request.
// TODO implement the decoder
func decodeGetPassengerInfoRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetPassengerInfoRequest)
	return endpoint.GetPassengerInfoRequest{
		Req: &pb.GetPassengerInfoRequest{
			Username: req.Username,
			Password: req.Password,
		},
	}, nil
}

// encodeGetPassengerInfoResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetPassengerInfoResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetPassengerInfoResponse)
	return &pb.GetPassengerInfoReply{
		UserId:               resp.Resp.UserId,
		Name:                 resp.Resp.Name,
		Age:                  resp.Resp.Age,
		AccountNum:           resp.Resp.AccountNum,
		Asset:                resp.Resp.Asset,
	}, resp.Err
}
func (g *grpcServer) GetPassengerInfo(ctx context1.Context, req *pb.GetPassengerInfoRequest) (*pb.GetPassengerInfoReply, error) {
	_, rep, err := g.getPassengerInfo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetPassengerInfoReply), nil
}

// makePublishOrderHandler creates the handler logic
func makePublishOrderHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.PublishOrderEndpoint, decodePublishOrderRequest, encodePublishOrderResponse, options...)
}

// decodePublishOrderResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain PublishOrder request.
// TODO implement the decoder
func decodePublishOrderRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PublishOrderRequest)
	return endpoint.PublishOrderRequest{
		Req: &pb.PublishOrderRequest{
			PassengerId:   req.GetPassengerId(),
			StartTime:     req.GetStartTime(),
			Origin:        req.GetOrigin(),
			Destination:   req.GetDestination(),
			PassengerName: req.GetPassengerName(),
		},
	}, nil
}

// encodePublishOrderResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodePublishOrderResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.PublishOrderResponse)
	return &pb.PublishOrderReply{
		Status:               resp.Resp.Status,
		DriverName:           resp.Resp.DriverName,
		Location:             resp.Resp.Location,
		Car:                  resp.Resp.Car,
		Path:                 resp.Resp.Path,
	}, resp.Err
}
func (g *grpcServer) PublishOrder(ctx context1.Context, req *pb.PublishOrderRequest) (*pb.PublishOrderReply, error) {
	_, rep, err := g.publishOrder.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PublishOrderReply), nil
}
