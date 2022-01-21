package grpc

import (
	"context"
	endpoint "payment/pkg/endpoint"
	pb "pkg/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

func makePayHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.PayEndpoint, decodePayRequest, encodePayResponse, options...)
}

func decodePayRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PayRequest)
	return endpoint.PayRequest{
		BillNum:     req.BillNum,
		AccountNum:  req.AccountNum,
		PayPassword: req.PayPassword,
	}, nil
}

func encodePayResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.PayResponse)
	if resp.Err != nil {
		return &pb.PayReply{
			Status: false,
			Msg:    resp.Err.Error(),
		}, nil
	}
	return &pb.PayReply{
		Status: true,
		Msg:    resp.Msg,
	}, nil
}
func (g *grpcServer) Pay(ctx context1.Context, req *pb.PayRequest) (*pb.PayReply, error) {
	_, rep, err := g.pay.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PayReply), nil
}
