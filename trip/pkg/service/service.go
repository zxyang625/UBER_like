package service

import (
	"context"
	"encoding/json"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"pkg/dao/mq"
	Err "pkg/error"
	"pkg/pb"
	"time"
)

// TripService describes the service.
type TripService interface {
	GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error)
}

type basicTripService struct{}

var defaultBasicTripService = &basicTripService{}

func (b *basicTripService) GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error) {
	traceID, _ := model.TraceIDFromHex(ctx.Value("Trace-ID").(string))
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	asyncReq := mq.AsyncReq{
		Method:        http.MethodPost,
		OriginApp:     "trip",
		OriginService: "gen-trip",
		DestApp:       "driver",
		DestService:   "take-order",
		TraceID:       traceID,
		Priority:      ctx.Value("Length").(int),
		Header:        nil,
		Data:          data,
	}
	mqData, err := json.Marshal(asyncReq)
	if err != nil {
		return nil, err
	}
	err = TripMessageServer.Publish(ctx, PublishQueueName, ctx.Value("Length").(int), mqData)
	rsp := &pb.TakeOrderReply{}
	c := make(chan struct{}, 1)
	d := amqp.Delivery{}
	go func() {
		d, err = TripMessageServer.ReceiveResp(ctx)
		c <- struct{}{}
	}()
	select {
	case <-c:
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(d.Body, rsp)
		if err != nil {
			log.Println("json.Unmarshal(d.Body, rsp) error ", err)
			return &pb.GenTripReply{Status: false, Msg: err.Error()}, err
		}
		return &pb.GenTripReply{Status: rsp.Status, Msg: rsp.Msg}, nil
	case <-time.After(time.Second):
		return &pb.GenTripReply{Status: false, Msg: "request timeout"}, Err.New(Err.RPCRequestTimeout, "PublishOrder timeout")
	}
}

// NewBasicTripService returns a naive, stateless implementation of TripService.
func NewBasicTripService() TripService {
	return &basicTripService{}
}

// New returns a TripService with all of the expected middleware wired in.
func New(middleware []Middleware) TripService {
	var svc TripService = NewBasicTripService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

//
//func RecvReqAndPushTrip(ctx context.Context, logger log.Logger, tracer *zipkin.Tracer) {
//	go func(*zipkin.Tracer) {
//		deliverServer := mq.InitDeliverMiddleware(tracer, "trip")(mq.HandleFunc(func(ctx context.Context, d amqp.Delivery) {
//			mqModel := &mq.MQModel{}
//			err := json.Unmarshal(d.Body, mqModel)
//			if err != nil {
//				logger.Log("method", "json.Unmarshal", "err", err)
//				return
//			}
//			r1 := &pb.PublishOrderRequest{}
//			err = proto.Unmarshal(mqModel.Data, r1)
//			if err != nil {
//				logger.Log("consume", ConsumePassengerName, "err", err)
//				return
//			}
//			err = redis.Passenger{}.LPush(r1)
//			if err != nil {
//				logger.Log("method", "LPUSH", "name", "passenger_list", "err", err)
//				return
//			}
//			d.Ack(false)
//			logger.Log("method", "consume", "name", ConsumePassengerName, "err", "null")
//
//			resp1 := &pb.PublishOrderReply{
//				Status: true,
//				Msg:    "publish order success! waiting in line...",
//			}
//			data, err := proto.Marshal(resp1)
//			if err != nil {
//				logger.Log("method", "proto marshal", "target", "PublishOrderReply", "err", err)
//				return
//			}
//			err = PassengerRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, data)
//			if err != nil {
//				logger.Log("method", "SendResp", "err", err)
//				return
//			}
//		}))
//		TripMessageServer.Consume(ctx, ConsumePassengerName, deliverServer)
//	}(tracer)
//
//	go func(context.Context) {
//		deliverServer := mq.InitDeliverMiddleware(tracer, "trip")(mq.HandleFunc(func(ctx context.Context, d amqp.Delivery) {
//			mqModel := &mq.MQModel{}
//			err := json.Unmarshal(d.Body, mqModel)
//			if err != nil {
//				logger.Log("method", "json.Unmarshal", "err", err)
//				return
//			}
//			r2 := &pb.TakeOrderRequest{}
//			err = proto.Unmarshal(mqModel.Data, r2)
//			if err != nil {
//				logger.Log("consume", ConsumeDriverName, "err", err)
//				return
//			}
//			err = redis.Driver{}.LPush(r2)
//			if err != nil {
//				logger.Log("method", "LPUSH", "name", "driver_list", "err", err)
//				return
//			}
//			d.Ack(false)
//			logger.Log("method", "consume", "name", ConsumeDriverName, "err", "null")
//
//			resp2 := &pb.TakeOrderReply{
//				Status: true,
//				Msg:    "take order success! waiting in line...",
//			}
//			data, err := proto.Marshal(resp2)
//			if err != nil {
//				logger.Log("method", "proto marshal", "target", "TakeOrderReply", "err", err)
//				return
//			}
//			err = DriverRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, data)
//			if err != nil {
//				logger.Log("method", "SendResp", "err", err)
//				return
//			}
//		}))
//		TripMessageServer.Consume(ctx, ConsumeDriverName, deliverServer)
//	}(ctx)
//
//	for {
//		passenger, err := redis.Passenger{}.BRPOP()
//		if err != nil {
//			logger.Log("method", "BRPOPData", "name", "passenger_list", "err", err)
//		}
//		driver, err := redis.Driver{}.BRPOP()
//		if err != nil {
//			logger.Log("method", "BRPOPData", "name", "driver_list", "err", err)
//		}
//
//		trip := &pb.TripMsg{
//			TripNum:       0, //自增，设置为0
//			PassengerId:   passenger.PassengerId,
//			DriverId:      driver.DriverId,
//			PassengerName: passenger.PassengerName,
//			DriverName:    driver.DriverName,
//			StartTime:     passenger.StartTime,
//			EndTime:       passenger.StartTime + rand.Int63n(3600), //随机
//			Origin:        passenger.Origin,
//			Destination:   passenger.Destination,
//			Car:           driver.Car,
//			Path:          "直走2km",
//		}
//		err = redis.Trip{}.LPUSH(trip)
//		if err != nil {
//			logger.Log("method", "LPUSH", "trip_num", trip.TripNum, "err", err)
//		}
//
//		span, ctx1 := tracer.StartSpanFromContext(ctx, "trip/gen_trip")
//		_, err = defaultBasicTripService.GenTrip(ctx1, &pb.GenTripRequest{})
//		if err != nil {
//			logger.Log("method", "GenTrip", "err", err)
//		}
//		span.Finish()
//	}
//}
