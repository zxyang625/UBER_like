package service

import (
	"billing"
	"context"
)

// BillingService describes the service.
type BillingService interface {
	GenBill(ctx context.Context, req *billing.GenBillRequest) (resp *billing.GenBillReply, err error)
	GetBillList(ctx context.Context, userId int64) (resp []*billing.BillMsg, err error)
}

type basicBillingService struct{}

func (b *basicBillingService) GenBill(ctx context.Context, req *billing.GenBillRequest) (resp *billing.GenBillReply, err error) {
	return &billing.GenBillReply{
		Status: true,
		BillMsg: &billing.BillMsg{
			BillNum:       1234,
			Price:         123.512,
			StartTime:     10,
			EndTime:       15,
			Origin:        "三元里",
			Destination:   "成华大道",
			PassengerName: "张三",
			DriverName:    "李四",
			Payed:         false,
		},
	}, nil
}

func (b *basicBillingService) GetBillList(ctx context.Context, userId int64) (resp []*billing.BillMsg, err error) {
	return []*billing.BillMsg{
		{
			BillNum: 1,
		},
		{
			BillNum: 2,
		},
		{
			BillNum: 3,
		},
	}, nil
}

// NewBasicBillingService returns a naive, stateless implementation of BillingService.
func NewBasicBillingService() BillingService {
	return &basicBillingService{}
}

// New returns a BillingService with all of the expected middleware wired in.
func New(middleware []Middleware) BillingService {
	var svc BillingService = NewBasicBillingService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
