package service

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
	"math/rand"
	"pkg/dao/models"
	"pkg/dao/redis"
	Err "pkg/error"
	"pkg/pb"
)

// TripService describes the service.
type TripService interface {
	GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error)
}

type basicTripService struct{}

var defaultBasicTripService = &basicTripService{}

func (b *basicTripService) GenTrip(ctx context.Context, req *pb.GenTripRequest) (resp *pb.GenTripReply, err error) {
	tripData, err := redis.Trip{}.BRPOPData()
	if err != nil {
		return nil, err
	}
	trip := &pb.TripMsg{}
	err = proto.Unmarshal(tripData, trip)
	if err != nil {
		return nil, err
	}
	trip1 := &models.Trip{
		TripNum:       trip.TripNum,
		PassengerId:   trip.PassengerId,
		DriverId:      trip.DriverId,
		PassengerName: trip.PassengerName,
		DriverName:    trip.DriverName,
		StartTime:     trip.StartTime,
		EndTime:       trip.EndTime,
		Origin:        trip.Origin,
		Destination:   trip.Destination,
		Car:           trip.Car,
		Path:          trip.Path,
	}
	err = models.AddTrip(trip1)
	if err != nil {
		return nil, err
	}
	err = TripMessageServer.Publish(ctx, PublishQueueName, tripData)
	if err != nil {
		return nil, Err.New(Err.ProtoUnmarshalFail, err.Error())
	}
	return &pb.GenTripReply{Status: true, Msg: "行程生成成功,正在生成订单..."}, nil
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

func RecvReqAndPushTrip(ctx context.Context, logger log.Logger) {
	go func(context.Context) {
		TripMessageServer.Consume(ctx, ConsumePassengerName, func(d amqp.Delivery) {
			r1 := &pb.PublishOrderRequest{}
			err := proto.Unmarshal(d.Body, r1)
			if err != nil {
				logger.Log("consume", ConsumePassengerName, "err", err)
				return
			}
			err = redis.Passenger{}.LPush(r1)
			if err != nil {
				logger.Log("method", "LPUSH", "name", "passenger_list", "err", err)
				return
			}
			d.Ack(false)
			logger.Log("method", "consume", "name", ConsumePassengerName, "err", "null")

			//go func() {
				resp1 := &pb.PublishOrderReply{
					Status: true,
					Msg: "publish order success! waiting in line...",
				}
				data, err := proto.Marshal(resp1)
				if err != nil {
					logger.Log("method", "proto marshal", "target", "PublishOrderReply", "err", err)
					return
				}
				err = PassengerRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, data)
				if err != nil {
					logger.Log("method", "SendResp", "err", err)
					return
				}
			//}()
		})
	}(ctx)

	go func(context.Context) {
		TripMessageServer.Consume(ctx, ConsumeDriverName, func(d amqp.Delivery) {
			r2 := &pb.TakeOrderRequest{}
			err := proto.Unmarshal(d.Body, r2)
			if err != nil {
				logger.Log("consume", ConsumeDriverName, "err", err)
				return
			}
			err = redis.Driver{}.LPush(r2)
			if err != nil {
				logger.Log("method", "LPUSH", "name", "driver_list", "err", err)
				return
			}
			d.Ack(false)
			logger.Log("method", "consume", "name", ConsumeDriverName, "err", "null")
			//go func() {
				resp2 := &pb.TakeOrderReply{
					Status: true,
					Msg: "take order success! waiting in line...",
				}
				data, err := proto.Marshal(resp2)
				if err != nil {
					logger.Log("method", "proto marshal", "target", "TakeOrderReply", "err", err)
					return
				}
				err = DriverRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, data)
				if err != nil {
					logger.Log("method", "SendResp", "err", err)
					return
				}
			//}()
		})
	}(ctx)

	for {
		passenger, err := redis.Passenger{}.BRPOP()
		if err != nil {
			logger.Log("method", "BRPOPData", "name", "passenger_list", "err", err)
		}
		driver, err := redis.Driver{}.BRPOP()
		if err != nil {
			logger.Log("method", "BRPOPData", "name", "driver_list", "err", err)
		}

		trip := &pb.TripMsg{
			TripNum:       0, //自增，设置为0
			PassengerId:   passenger.PassengerId,
			DriverId:      driver.DriverId,
			PassengerName: passenger.PassengerName,
			DriverName:    driver.DriverName,
			StartTime:     passenger.StartTime,
			EndTime:       passenger.StartTime + rand.Int63n(3600), //随机
			Origin:        passenger.Origin,
			Destination:   passenger.Destination,
			Car:           driver.Car,
			Path:          "直走2km",
		}
		err = redis.Trip{}.LPUSH(trip)
		if err != nil {
			logger.Log("method", "LPUSH", "trip_num", trip.TripNum, "err", err)
		}

		_, err = defaultBasicTripService.GenTrip(ctx, &pb.GenTripRequest{})
		if err != nil {
			logger.Log("method", "GenTrip", "err", err)
		}
	}
}
