package grpc

import (
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	"trip"
	endpoint "trip/pkg/endpoint"
	pb "trip/pkg/grpc/pb"
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
		Req: &trip.GenTripRequest{
			PassengerReq: &trip.PublishOrderRequest{
				PassengerId:   req.PassengerReq.GetPassengerId(),
				StartTime:     req.PassengerReq.GetStartTime(),
				Origin:        req.PassengerReq.GetOrigin(),
				Destination:   req.PassengerReq.GetDestination(),
				PassengerName: req.PassengerReq.GetPassengerName(),
			},
			DriverReq:    &trip.TakeOrderRequest{
				DriverId:   req.DriverReq.GetDriverId(),
				DriverName: req.DriverReq.GetDriverName(),
				Location:   req.DriverReq.GetLocation(),
				Car:        req.DriverReq.GetCar(),
			},
		},
	}, nil
}

// encodeGenTripResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGenTripResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenTripResponse)
	return &pb.GenTripReply{
		Status:      resp.Resp.Status,
		TripMsg:     &pb.TripMsg{
			TripNum:              resp.Resp.Trip.TripNum,
			PassengerId:          resp.Resp.Trip.PassengerId,
			DriverId:             resp.Resp.Trip.DriverId,
			PassengerName:        resp.Resp.Trip.PassengerName,
			DriverName:           resp.Resp.Trip.DriverName,
			StartTime:            resp.Resp.Trip.StartTime,
			EndTime:              resp.Resp.Trip.EndTime,
			Origin:               resp.Resp.Trip.Origin,
			Destination:          resp.Resp.Trip.Destination,
			Car:                  resp.Resp.Trip.Car,
			Path:                 resp.Resp.Trip.Path,
		},
	}, resp.Err
}
func (g *grpcServer) GenTrip(ctx context1.Context, req *pb.GenTripRequest) (*pb.GenTripReply, error) {
	_, rep, err := g.genTrip.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenTripReply), nil
}
