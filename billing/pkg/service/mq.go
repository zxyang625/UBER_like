package service

import (
	"pkg/dao/mq"
)

var (
	BillingMessageServer mq.MessageServer
	TripRespMessageServer mq.MessageServer
)

const (
	PublishQueueName = ""
	ConsumeQueueName = "trip_queue"
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	BillingMessageServer, err = mq.NewMessageServer("billing")
	if err != nil {
		return err
	}
	TripRespMessageServer, err = mq.NewMessageServer("")
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		BillingMessageServer = mdw(BillingMessageServer)
		TripRespMessageServer = mdw(TripRespMessageServer)
	}
	return nil
}
