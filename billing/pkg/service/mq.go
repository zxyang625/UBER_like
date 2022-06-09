package service

import (
	"pkg/dao/mq"
)

var (
	BillingMessageServer mq.MessageServer
	PayRespMessageServer mq.MessageServer
	//TripRespMessageServer mq.MessageServer
)

const (
	ConsumeTripQueueName = "trip_queue"
	ConsumePayQueueName  = "pay_queue"
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	BillingMessageServer, err = mq.NewMessageServer("billing_queue", 5)
	if err != nil {
		return err
	}
	//TripRespMessageServer, err = mq.NewMessageServer("")
	//if err != nil {
	//	return err
	//}
	for _, mdw := range mdws {
		BillingMessageServer = mdw(BillingMessageServer)
		//TripRespMessageServer = mdw(TripRespMessageServer)
	}
	return nil
}

func InitPaySendRespMessageServer(mdws ...mq.Middleware) error {
	var err error
	PayRespMessageServer, err = mq.NewMessageServer("payment_reply", 3)
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		PayRespMessageServer = mdw(PayRespMessageServer)
	}
	return nil
}
