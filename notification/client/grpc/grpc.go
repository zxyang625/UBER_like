package grpc

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	endpoint1 "notification/pkg/endpoint"
	pb "notification/pkg/grpc/pb"
	service "notification/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.NotificationService, error) {
	var noticeTripEndpoint endpoint.Endpoint
	{
		noticeTripEndpoint = grpc1.NewClient(conn, "pb.Notification", "NoticeTrip", encodeNoticeTripRequest, decodeNoticeTripResponse, pb.NoticeTripReply{}, options["NoticeTrip"]...).Endpoint()
	}

	return endpoint1.Endpoints{NoticeTripEndpoint: noticeTripEndpoint}, nil
}

// encodeNoticeTripRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain NoticeTrip request to a gRPC request.
func encodeNoticeTripRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.NoticeTripRequest)
	data, err := proto.Marshal(req.Req)
	if err != nil {
		return nil, err
	}
	pbReq := &pb.NoticeTripRequest{}
	err = proto.Unmarshal(data, pbReq)
	if err != nil {
		return nil, err
	}
	return pbReq, nil
}

// decodeNoticeTripResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeNoticeTripResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.NoticeTripReply)
	data, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	endResp := endpoint1.NoticeTripResponse{Resp: new(pb.NoticeTripReply)}
	err = proto.Unmarshal(data, endResp.Resp)
	if err != nil {
		return nil, err
	}
	return endResp, nil
}
