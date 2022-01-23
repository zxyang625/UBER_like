package grpc

import (
	endpoint1 "billing/pkg/endpoint"
	service "billing/pkg/service"
	"context"
	"errors"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	pb "pkg/pb"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.BillingService, error) {
	var genBillEndpoint endpoint.Endpoint
	{
		genBillEndpoint = grpc1.NewClient(conn, "pb.Billing", "GenBill", encodeGenBillRequest, decodeGenBillResponse, pb.GenBillReply{}, options["GenBill"]...).Endpoint()
	}

	var getBillListEndpoint endpoint.Endpoint
	{
		getBillListEndpoint = grpc1.NewClient(conn, "pb.Billing", "GetBillList", encodeGetBillListRequest, decodeGetBillListResponse, pb.GetBillListReply{}, options["GetBillList"]...).Endpoint()
	}

	var getBillEndpoint endpoint.Endpoint
	{
		getBillEndpoint = grpc1.NewClient(conn, "pb.Billing", "GetBill", encodeGetBillRequest, decodeGetBillResponse, pb.GetBillReply{}, options["GetBill"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		GenBillEndpoint:     genBillEndpoint,
		GetBillEndpoint:     getBillEndpoint,
		GetBillListEndpoint: getBillListEndpoint,
	}, nil
}

// encodeGenBillRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenBill request to a gRPC request.
func encodeGenBillRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(endpoint1.GenBillRequest)
	return &pb.GenBillRequest{}, nil
}

// decodeGenBillResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenBillResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Billing' Decoder is not impelemented")
}

// encodeGetBillListRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetBillList request to a gRPC request.
func encodeGetBillListRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Billing' Encoder is not impelemented")
}

// decodeGetBillListResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetBillListResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Billing' Decoder is not impelemented")
}

// encodeGetBillRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetBill request to a gRPC request.
func encodeGetBillRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Billing' Encoder is not impelemented")
}

// decodeGetBillResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetBillResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Billing' Decoder is not impelemented")
}
