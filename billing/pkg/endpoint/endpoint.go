package endpoint

import (
	"pkg/pb"
	service "billing/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GenBillRequest collects the request parameters for the GenBill method.
type GenBillRequest struct {
	Req *pb.GenBillRequest `json:"req"`
}

// GenBillResponse collects the response parameters for the GenBill method.
type GenBillResponse struct {
	Resp *pb.GenBillReply `json:"resp"`
	Err  error            `json:"err"`
}

// MakeGenBillEndpoint returns an endpoint that invokes GenBill on the service.
func MakeGenBillEndpoint(s service.BillingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenBillRequest)
		resp, err := s.GenBill(ctx, req.Req)
		return GenBillResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GenBillResponse) Failed() error {
	return r.Err
}

// GetBillListRequest collects the request parameters for the GetBillList method.
type GetBillListRequest struct {
	UserId int64 `json:"user_id"`
}

// GetBillListResponse collects the response parameters for the GetBillList method.
type GetBillListResponse struct {
	Resp []*pb.BillMsg `json:"resp"`
	Err  error         `json:"err"`
}

// MakeGetBillListEndpoint returns an endpoint that invokes GetBillList on the service.
func MakeGetBillListEndpoint(s service.BillingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBillListRequest)
		resp, err := s.GetBillList(ctx, req.UserId)
		return GetBillListResponse{
			Err:  err,
			Resp: resp,
		}, nil
	}
}

// Failed implements Failer.
func (r GetBillListResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GenBill implements Service. Primarily useful in a client.
func (e Endpoints) GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error) {
	request := GenBillRequest{Req: req}
	response, err := e.GenBillEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GenBillResponse).Resp, response.(GenBillResponse).Err
}

// GetBillList implements Service. Primarily useful in a client.
func (e Endpoints) GetBillList(ctx context.Context, userId int64) (resp []*pb.BillMsg, err error) {
	request := GetBillListRequest{UserId: userId}
	response, err := e.GetBillListEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetBillListResponse).Resp, response.(GetBillListResponse).Err
}
