package grpc

import (
	endpoint "billing/pkg/endpoint"
	"context"
	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	context1 "golang.org/x/net/context"
	"pkg/pb"
)

// makeGenBillHandler creates the handler logic
func makeGenBillHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenBillEndpoint, decodeGenBillRequest, encodeGenBillResponse, options...)
}

// decodeGenBillResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GenBill request.
// TODO implement the decoder
func decodeGenBillRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GenBillRequest)
	return endpoint.GenBillRequest{
		Req: &pb.GenBillRequest{
			TripMsg: &pb.TripMsg{
				TripNum:       req.TripMsg.TripNum,
				PassengerId:   req.TripMsg.PassengerId,
				DriverId:      req.TripMsg.DriverId,
				PassengerName: req.TripMsg.PassengerName,
				DriverName:    req.TripMsg.DriverName,
				StartTime:     req.TripMsg.StartTime,
				EndTime:       req.TripMsg.EndTime,
				Origin:        req.TripMsg.Origin,
				Destination:   req.TripMsg.Destination,
				Car:           req.TripMsg.Car,
				Path:          req.TripMsg.Path,
			},
		},
	}, nil
}

// encodeGenBillResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGenBillResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenBillResponse)
	return &pb.GenBillReply{
		Status:               resp.Resp.Status,
		BillMsg:              &pb.BillMsg{
			BillNum:              resp.Resp.BillMsg.BillNum,
			Price:                resp.Resp.BillMsg.Price,
			StartTime:            resp.Resp.BillMsg.StartTime,
			EndTime:              resp.Resp.BillMsg.EndTime,
			Origin:               resp.Resp.BillMsg.Origin,
			Destination:          resp.Resp.BillMsg.Destination,
			PassengerName:        resp.Resp.BillMsg.PassengerName,
			DriverName:           resp.Resp.BillMsg.DriverName,
			Payed:                resp.Resp.BillMsg.Payed,
		},
	}, resp.Err
}
func (g *grpcServer) GenBill(ctx context1.Context, req *pb.GenBillRequest) (*pb.GenBillReply, error) {
	_, rep, err := g.genBill.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GenBillReply), nil
}

// makeGetBillListHandler creates the handler logic
func makeGetBillListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetBillListEndpoint, decodeGetBillListRequest, encodeGetBillListResponse, options...)
}

// decodeGetBillListResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetBillList request.
// TODO implement the decoder
func decodeGetBillListRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetBillListRequest)
	return endpoint.GetBillListRequest{UserId: req.GetUserId()}, nil
}

// encodeGetBillListResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetBillListResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetBillListResponse)
	pbResp := &pb.GetBillListReply{
		BillList: make([]*pb.BillMsg, 0, len(resp.Resp)),
	}
	for i := range resp.Resp {
		data, err := proto.Marshal(resp.Resp[i])
		if err != nil {
			return nil, err
		}
		temp := &pb.BillMsg{}
		err = proto.Unmarshal(data, temp)
		if err != nil {
			return nil, err
		}
		pbResp.BillList = append(pbResp.BillList, temp)
	}
	return pbResp, resp.Err
}
func (g *grpcServer) GetBillList(ctx context1.Context, req *pb.GetBillListRequest) (*pb.GetBillListReply, error) {
	_, rep, err := g.getBillList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetBillListReply), nil
}
