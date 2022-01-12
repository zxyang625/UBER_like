package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	context1 "golang.org/x/net/context"
	endpoint "notification/pkg/endpoint"
	"notification/pkg/grpc/pb"
)

// makeNoticeTripHandler creates the handler logic
func makeNoticeTripHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.NoticeTripEndpoint, decodeNoticeTripRequest, encodeNoticeTripResponse, options...)
}

// decodeNoticeTripResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain NoticeTrip request.
// TODO implement the decoder
func decodeNoticeTripRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.NoticeTripRequest)
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	endReq := endpoint.NoticeTripRequest{Req: new(pb.NoticeTripRequest)}
	err = proto.Unmarshal(data, endReq.Req)
	if err != nil {
		return nil, err
	}
	return endReq, nil
}

// encodeNoticeTripResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeNoticeTripResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.NoticeTripResponse)
	data, err := proto.Marshal(resp.Resp)
	if err != nil {
		return nil, err
	}
	pbResp := &pb.NoticeTripReply{}
	err = proto.Unmarshal(data, pbResp)
	if err != nil {
		return nil, err
	}
	return pbResp, nil
}

func (g *grpcServer) NoticeTrip(ctx context1.Context, req *pb.NoticeTripRequest) (*pb.NoticeTripReply, error) {
	_, rep, err := g.noticeTrip.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.NoticeTripReply), nil
}