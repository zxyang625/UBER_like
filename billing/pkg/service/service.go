package service

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/proto"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"math"
	"math/rand"
	"net/http"
	"pkg/dao/models"
	"pkg/dao/mq"
	"pkg/dao/redis"
	"pkg/pb"
)

// BillingService describes the service.
type BillingService interface {
	GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error)
	GetBillList(ctx context.Context, userId int64) (resp []*pb.BillMsg, err error)
	GetBill(ctx context.Context, billNum int64) (resp *pb.BillMsg, err error)
	SetPayedAndGetPrice(ctx context.Context, billNum int64) (float32, error)
}

type basicBillingService struct{}

var defaultBasicBillingService = &basicBillingService{}

func (b *basicBillingService) SetPayedAndGetPrice(ctx context.Context, billNum int64) (float32, error) {
	price, err := models.SetPayedAndGetPrice(billNum)
	if err != nil {
		return 0, fmt.Errorf("SetPayedAndGetPrice failed, err: %v", err)
	}
	return price, err
}

func (b *basicBillingService) GenBill(ctx context.Context, req *pb.GenBillRequest) (resp *pb.GenBillReply, err error) {
	bill1 := &models.Bill{
		BillNum:       0,
		Price:         rand.Float32(),
		StartTime:     req.BillMsg.GetStartTime(),
		EndTime:       req.BillMsg.GetEndTime(),
		Origin:        req.BillMsg.GetOrigin(),
		Destination:   req.BillMsg.GetDestination(),
		PassengerName: req.BillMsg.GetPassengerName(),
		DriverName:    req.BillMsg.GetDriverName(),
		Payed:         false,
		PassengerId:   req.BillMsg.GetPassengerId(),
		DriverId:      req.BillMsg.GetDriverId(),
	}
	err = models.AddBill(bill1)
	if err != nil {
		return &pb.GenBillReply{Status: false}, err
	}
	return nil, nil
}

func (b *basicBillingService) GetBill(ctx context.Context, billNum int64) (resp *pb.BillMsg, err error) {
	traceID, _ := model.TraceIDFromHex(ctx.Value("Trace-ID").(string))
	p := ctx.Value("Priority").(int)
	if p >= 5 {
		return nil, nil
	}
	//bill, err := models.GetBill(billNum)
	//if err != nil {
	//	return nil, err
	//}
	//data, err := json.Marshal(&models.Bill{BillNum: 250205, Price: 10, Origin: "?????????", Destination: "????????????"})
	req := mq.AsyncReq{
		Method:        http.MethodGet,
		OriginApp:     "billing",
		OriginService: "get-bill",
		DestApp:       "billing",
		DestService:   "get-bill",
		TraceID:       traceID,
		Length:        ctx.Value("Length").(int),
		Header:        nil,
		Data:          nil,
	}
	if err != nil {
		return nil, err
	}
	mqData, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	/////////////////////////////////////////////
	//if 1 <= ctx.Value("Priority").(int) && ctx.Value("Priority").(int) <= 5 {
	err = BillingMessageServer.Publish(ctx, "billing_queue", ctx.Value("Priority").(int), mqData)
	//} else if ctx.Value("Priority").(int) == 0 {
	//	err = BillingMessageServer.Publish(ctx, "billing_queue", 1, mqData)
	//}
	resp = &pb.BillMsg{}
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

func RecvAndGenBill(ctx context.Context, logger log.Logger, tracer *zipkin.Tracer) {
	go func(*zipkin.Tracer) {
		deliverServer := mq.InitDeliverMiddleware(tracer, "billing")(mq.HandleFunc(func(ctx context.Context, d amqp.Delivery) {
			mqModel := &mq.MQModel{}
			err := json.Unmarshal(d.Body, mqModel)
			if err != nil {
				logger.Log("mehtod", "json.Unmarshal", "err", err)
				return
			}
			r := &pb.TripMsg{}
			err = proto.Unmarshal(mqModel.Data, r)
			if err != nil {
				logger.Log("consume", ConsumeTripQueueName, "err", err)
				return
			}
			d.Ack(false)
			logger.Log("method", "consume", "name", ConsumeTripQueueName, "err", "null")

			err = redis.Billing{}.LPUSH(&pb.BillMsg{
				BillNum:       0,
				Price:         rand.Float32(), //??????????????????
				StartTime:     r.GetStartTime(),
				EndTime:       r.GetEndTime(),
				Origin:        r.GetOrigin(),
				Destination:   r.GetDestination(),
				PassengerName: r.GetPassengerName(),
				DriverName:    r.GetDriverName(),
				Payed:         false,
				PassengerId:   r.GetPassengerId(),
				DriverId:      r.GetDriverId(),
			})
			if err != nil {
				logger.Log("method", "LPUSH", "name", "bill_list", "err", err)
				return
			}
		}))
		BillingMessageServer.Consume(ctx, ConsumeTripQueueName, deliverServer)
	}(tracer)

	go func(tracer *zipkin.Tracer) {
		deliverServer := mq.InitDeliverMiddleware(tracer, "billing")(mq.HandleFunc(func(ctx context.Context, d amqp.Delivery) {
			d.Ack(false)
			mqModel := &mq.MQModel{}
			err := json.Unmarshal(d.Body, mqModel)
			if err != nil {
				logger.Log("mehtod", "json.Unmarshal", "err", err)
				return
			}
			billNum := int64(binary.BigEndian.Uint64(mqModel.Data))
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
		}))
		BillingMessageServer.Consume(ctx, ConsumePayQueueName, deliverServer)
	}(tracer)

	for {
		billData, err := redis.Billing{}.BRPOPData()
		if err != nil {
			logger.Log("method", "BRPOPData", "name", "bill_list", "err", err)
		}
		bill := &pb.BillMsg{}
		err = proto.Unmarshal(billData, bill)
		if err != nil {
			logger.Log("method", "proto.Unmarshal", "err", err)
		}
		_, err = defaultBasicBillingService.GenBill(ctx, &pb.GenBillRequest{
			BillMsg: bill,
		})
		if err != nil {
			logger.Log("method", "GenBill", "err", err)
		}
	}
}
