package grpc

import (
	"context"
	"errors"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
	pb "pkg/pb"
	endpoint1 "trip/pkg/endpoint"
	service "trip/pkg/service"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.TripService, error) {
	var genTripEndpoint endpoint.Endpoint
	{
		genTripEndpoint = grpc1.NewClient(conn, "pb.Trip", "GenTrip", encodeGenTripRequest, decodeGenTripResponse, pb.GenTripReply{}, options["GenTrip"]...).Endpoint()
	}

	return endpoint1.Endpoints{GenTripEndpoint: genTripEndpoint}, nil
}

// encodeGenTripRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenTrip request to a gRPC request.
func encodeGenTripRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Trip' Encoder is not impelemented")
}

// decodeGenTripResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenTripResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Trip' Decoder is not impelemented")
}
