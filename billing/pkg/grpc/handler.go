package grpc

import (
	endpoint "billing/pkg/endpoint"
	"context"
	"pkg/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	context1 "golang.org/x/net/context"
)

func makeGenBillHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GenBillEndpoint, decodeGenBillRequest, encodeGenBillResponse, options...)
}

func decodeGenBillRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GenBillRequest)
	return endpoint.GenBillRequest{
		Req: &pb.GenBillRequest{
			BillMsg: &pb.BillMsg{
				BillNum:       req.BillMsg.BillNum,
				PassengerId:   req.BillMsg.PassengerId,
				DriverId:      req.BillMsg.DriverId,
				PassengerName: req.BillMsg.PassengerName,
				DriverName:    req.BillMsg.DriverName,
				StartTime:     req.BillMsg.StartTime,
				EndTime:       req.BillMsg.EndTime,
				Origin:        req.BillMsg.Origin,
				Destination:   req.BillMsg.Destination,
			},
		},
	}, nil
}

func encodeGenBillResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GenBillResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.GenBillReply{
		Status: resp.Resp.Status,
		BillMsg: &pb.BillMsg{
			BillNum:       resp.Resp.BillMsg.BillNum,
			Price:         resp.Resp.BillMsg.Price,
			StartTime:     resp.Resp.BillMsg.StartTime,
			EndTime:       resp.Resp.BillMsg.EndTime,
			Origin:        resp.Resp.BillMsg.Origin,
			Destination:   resp.Resp.BillMsg.Destination,
			PassengerName: resp.Resp.BillMsg.PassengerName,
			DriverName:    resp.Resp.BillMsg.DriverName,
			Payed:         resp.Resp.BillMsg.Payed,
			PassengerId:   resp.Resp.BillMsg.PassengerId,
			DriverId:      resp.Resp.BillMsg.DriverId,
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

func makeGetBillListHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetBillListEndpoint, decodeGetBillListRequest, encodeGetBillListResponse, options...)
}

func decodeGetBillListRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetBillListRequest)
	return endpoint.GetBillListRequest{UserId: req.GetUserId()}, nil
}

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

func makeGetBillHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetBillEndpoint, decodeGetBillRequest, encodeGetBillResponse, options...)
}

func decodeGetBillRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetBillRequest)
	return endpoint.GetBillRequest{
		BillNum: req.GetBillNum(),
	}, nil
}

func encodeGetBillResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(endpoint.GetBillResponse)
	if resp.Err != nil {
		return &pb.GetBillReply{Status: false}, resp.Err
	}
	return &pb.GetBillReply{
		Status:               true,
		BillMsg:              resp.Resp,
	}, nil
}
func (g *grpcServer) GetBill(ctx context1.Context, req *pb.GetBillRequest) (*pb.GetBillReply, error) {
	_, rep, err := g.getBill.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetBillReply), nil
}
