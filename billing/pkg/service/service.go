package service

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"math"
	"math/rand"
	"pkg/dao/models"
	"pkg/dao/redis"
	"pkg/pb"
)

// BillingService describes the service.
type BillingService interface {
	GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error)
	GetBillList(ctx context.Context, userId int64) (resp []*pb.BillMsg, err error)
	GetBill(ctx context.Context, billNum int64) (resp *pb.BillMsg, err error)
}

type basicBillingService struct{}

var defaultBasicBillingService = &basicBillingService{}

func (b *basicBillingService) GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error) {
	billData, err := redis.Billing{}.BRPOPData()
	if err != nil {
		return nil, err
	}
	bill := &pb.BillMsg{}
	err = proto.Unmarshal(billData, bill)
	if err != nil {
		return nil, err
	}
	bill1 := &models.Bill{
		BillNum:       0,
		Price:         rand.Float32(),
		StartTime:     bill.GetStartTime(),
		EndTime:       bill.GetEndTime(),
		Origin:        bill.GetOrigin(),
		Destination:   bill.GetDestination(),
		PassengerName: bill.GetPassengerName(),
		DriverName:    bill.GetDriverName(),
		Payed:         false,
		PassengerId:   bill.GetPassengerId(),
		DriverId:      bill.GetDriverId(),
	}
	err = models.AddBill(bill1)
	if err != nil {
		return &pb.GenBillReply{Status: false}, err
	}
	return nil, nil
}

func (b *basicBillingService) GetBill(ctx context.Context, billNum int64) (resp *pb.BillMsg, err error) {
	bill, err := models.GetBill(billNum)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(bill)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
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

func RecvAndGenBill(ctx context.Context, logger log.Logger) {
	go func(context.Context) {
		BillingMessageServer.Consume(ctx, ConsumeTripQueueName, func(d amqp.Delivery) {
			r := &pb.TripMsg{}
			err := proto.Unmarshal(d.Body, r)
			if err != nil {
				logger.Log("consume", ConsumeTripQueueName, "err", err)
				return
			}
			err = redis.Billing{}.LPUSH(&pb.BillMsg{
				BillNum:              0,
				Price:                rand.Float32(),		//价格随机计算
				StartTime:            r.GetStartTime(),
				EndTime:              r.GetEndTime(),
				Origin:               r.GetOrigin(),
				Destination:          r.GetDestination(),
				PassengerName:        r.GetPassengerName(),
				DriverName:           r.GetDriverName(),
				Payed:                false,
				PassengerId:          r.GetPassengerId(),
				DriverId:             r.GetDriverId(),
			})
			if err != nil {
				logger.Log("method", "LPUSH", "name", "bill_list", "err", err)
				return
			}
			logger.Log("method", "consume", "name", ConsumeTripQueueName, "err", "null")
			d.Ack(false)

			//err = TripRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, []byte("send trip success"))
			//if err != nil {
			//	logger.Log("method", "SendResp", "err", err)
			//	return
			//}
		})
	}(ctx)

	go func(context.Context) {
		BillingMessageServer.Consume(ctx, ConsumePayQueueName, func(d amqp.Delivery) {
			billNum := int64(binary.BigEndian.Uint64(d.Body))
			price, err := models.SetPayedAndGetPrice(billNum)
			if err != nil {
				logger.Log("method", "SetPayedAndGetPrice", "err", err)
				return
			}
			bits := math.Float32bits(price)
			data := make([]byte, 4)
			binary.LittleEndian.PutUint32(data, bits)
			err = PayRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, data)
			if err != nil {
				logger.Log("method", "SendResp", "err", err)
				return
			}
		})
	}(ctx)

	for {
		_, err := defaultBasicBillingService.GenBill(ctx, &pb.GenBillRequest{})
		if err != nil {
			logger.Log("method", "GenBill", "err", err)
		}
	}
}