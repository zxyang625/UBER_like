package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	pb "pkg/pb"
	endpoint "trip/pkg/endpoint"
)

// makeGenTripHandler creates the handler logic
func makeGenTripHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenTripEndpoint, decodeGenTripRequest, encodeGenTripResponse, options...)
}

// decodeGenTripResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GenTrip request.
// TODO implement the decoder
func decodeGenTripRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GenTripRequest)
	return endpoint.GenTripRequest{
		Req: req,
	}, nil
}

// encodeGenTripResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGenTripResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenTripResponse)
	if resp.Resp == nil {
		return nil, resp.Err
	}
	return &pb.GenTripReply{
		Status: resp.Resp.Status,
		Msg:    resp.Resp.Msg,
		Trip:   resp.Resp.Trip,
	}, resp.Err
}
func (g *grpcServer) GenTrip(ctx context1.Context, req *pb.GenTripRequest) (*pb.GenTripReply, error) {
	_, rep, err := g.genTrip.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenTripReply), nil
}
