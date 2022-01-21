package grpc

import (
	"context"
	"errors"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	endpoint1 "payment/pkg/endpoint"
	service "payment/pkg/service"
	pb "pkg/pb"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.PaymentService, error) {
	var payEndpoint endpoint.Endpoint
	{
		payEndpoint = grpc1.NewClient(conn, "pb.Payment", "Pay", encodePayRequest, decodePayResponse, pb.PayReply{}, options["Pay"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		PayEndpoint:             payEndpoint,
	}, nil
}

// encodePayRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Pay request to a gRPC request.
func encodePayRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.PayRequest)
	return &pb.PayRequest{
		BillNum:              req.BillNum,
		AccountNum:           req.AccountNum,
		PayPassword:          req.PayPassword,
	}, nil
}

// decodePayResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodePayResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.PayReply)
	if resp.Status != false {
		return endpoint1.PayResponse{
			Msg: "success" + resp.Msg,
			Err: nil,
		}, nil
	}
	return endpoint1.PayResponse{
		Msg: "failed",
		Err: errors.New(resp.Msg),
	}, nil
}
