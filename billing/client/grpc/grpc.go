package grpc

import (
	"billing"
	endpoint1 "billing/pkg/endpoint"
	pb "billing/pkg/grpc/pb"
	service "billing/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
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

	return endpoint1.Endpoints{
		GenBillEndpoint:     genBillEndpoint,
		GetBillListEndpoint: getBillListEndpoint,
	}, nil
}

// encodeGenBillRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GenBill request to a gRPC request.
func encodeGenBillRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.GenBillRequest)
	return &pb.GenBillRequest{
		TripMsg: &pb.TripMsg{
			TripNum:              req.Req.TripMsg.TripNum,
			PassengerId:          req.Req.TripMsg.PassengerId,
			DriverId:             req.Req.TripMsg.DriverId,
			PassengerName:        req.Req.TripMsg.PassengerName,
			DriverName:           req.Req.TripMsg.DriverName,
			StartTime:            req.Req.TripMsg.StartTime,
			EndTime:              req.Req.TripMsg.EndTime,
			Origin:               req.Req.TripMsg.Origin,
			Destination:          req.Req.TripMsg.Destination,
			Car:                  req.Req.TripMsg.Car,
			Path:                 req.Req.TripMsg.Path,
		},
	}, nil
}

// decodeGenBillResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGenBillResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.GenBillReply)
	return endpoint1.GenBillResponse{
		Resp: &billing.GenBillReply{
			Status:  resp.GetStatus(),
			BillMsg: &billing.BillMsg{
				BillNum:       resp.GetBillMsg().GetBillNum(),
				Price:         resp.GetBillMsg().GetPrice(),
				StartTime:     resp.GetBillMsg().GetStartTime(),
				EndTime:       resp.GetBillMsg().GetEndTime(),
				Origin:        resp.GetBillMsg().GetOrigin(),
				Destination:   resp.GetBillMsg().GetDestination(),
				PassengerName: resp.GetBillMsg().GetPassengerName(),
				DriverName:    resp.GetBillMsg().GetDriverName(),
				Payed:         resp.GetBillMsg().GetPayed(),
			},
		},
		Err:  nil,
	}, nil
}

// encodeGetBillListRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetBillList request to a gRPC request.
func encodeGetBillListRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.GetBillListRequest)
	return &pb.GetBillListRequest{
		UserId:               req.UserId,
	}, nil
}

// decodeGetBillListResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetBillListResponse(_ context.Context, reply interface{}) (interface{}, error) {
	resp := reply.(*pb.GetBillListReply)
	endResp := endpoint1.GetBillListResponse{
		Resp: make([]*billing.BillMsg, 0, len(resp.BillMsgList)),
		Err:  nil,
	}
	for i := range resp.BillMsgList {
		endResp.Resp = append(endResp.Resp, &billing.BillMsg{
			BillNum:       resp.BillMsgList[i].GetBillNum(),
			Price:         resp.BillMsgList[i].GetPrice(),
			StartTime:     resp.BillMsgList[i].GetStartTime(),
			EndTime:       resp.BillMsgList[i].GetEndTime(),
			Origin:        resp.BillMsgList[i].GetOrigin(),
			Destination:   resp.BillMsgList[i].GetDestination(),
			PassengerName: resp.BillMsgList[i].GetPassengerName(),
			DriverName:    resp.BillMsgList[i].GetDriverName(),
			Payed:         resp.BillMsgList[i].GetPayed(),
		})
	}
	return endResp, nil
}
