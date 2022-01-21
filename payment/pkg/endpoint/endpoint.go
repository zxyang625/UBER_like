package endpoint

import (
	"context"
	service "payment/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// PayRequest collects the request parameters for the Pay method.
type PayRequest struct {
	BillNum     int64  `json:"bill_num"`
	AccountNum  int64  `json:"account_num"`
	PayPassword string `json:"pay_password"`
}

// PayResponse collects the response parameters for the Pay method.
type PayResponse struct {
	Msg string `json:"msg"`
	Err error  `json:"err"`
}

// MakePayEndpoint returns an endpoint that invokes Pay on the service.
func MakePayEndpoint(s service.PaymentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PayRequest)
		msg, err := s.Pay(ctx, req.BillNum, req.AccountNum, req.PayPassword)
		return PayResponse{
			Err: err,
			Msg: msg,
		}, nil
	}
}

// Failed implements Failer.
func (r PayResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Pay implements Service. Primarily useful in a client.
func (e Endpoints) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	request := PayRequest{
		AccountNum:  accountNum,
		BillNum:     billNum,
		PayPassword: payPassword,
	}
	response, err := e.PayEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PayResponse).Msg, response.(PayResponse).Err
}
