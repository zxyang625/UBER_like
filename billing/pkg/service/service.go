package service

import (
	"context"
	"encoding/json"
	"pkg/dao/models"
	"pkg/pb"
)

// BillingService describes the service.
type BillingService interface {
	GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error)
	GetBillList(ctx context.Context, userId int64) (resp []*pb.BillMsg, err error)
}

type basicBillingService struct{}

func (b *basicBillingService) GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error) {
	return &pb.GenBillReply{
		Status: true,
		BillMsg: &pb.BillMsg{
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

func (b *basicBillingService) GetBillList(ctx context.Context, userId int64) ([]*pb.BillMsg, error) {
	list, err := models.GetBillList(userId)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	var resp []*pb.BillMsg
	err = json.Unmarshal(data, &resp)
	return resp, nil
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
