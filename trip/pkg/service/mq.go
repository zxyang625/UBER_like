package service

import (
	"pkg/dao/mq"
)

var (
	TripMessageServer          mq.MessageServer
	PassengerRespMessageServer mq.MessageServer
	DriverRespMessageServer    mq.MessageServer
)

const (
	PublishQueueName     = "trip_queue"
	ConsumePassengerName = "passenger_queue"
	ConsumeDriverName    = "driver_queue"
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	TripMessageServer, err = mq.NewMessageServer("trip_queue", 3)
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		TripMessageServer = mdw(TripMessageServer)
	}
	return nil
}

func InitTripSendRespMessageServer(mdws ...mq.Middleware) error {
	var err error
	PassengerRespMessageServer, err = mq.NewMessageServer("", 3)
	if err != nil {
		return err
	}
	DriverRespMessageServer, err = mq.NewMessageServer("", 3)
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		PassengerRespMessageServer = mdw(PassengerRespMessageServer)
		DriverRespMessageServer = mdw(DriverRespMessageServer)
	}
	return nil
}
