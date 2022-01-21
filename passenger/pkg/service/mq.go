package service

import (
	"pkg/dao/mq"
)

var PassengerMessageServer mq.MessageServer

const (
	PublishQueueName = "passenger_queue"
	ConsumeQueueName = "notification_queue"
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	PassengerMessageServer, err = mq.NewMessageServer( "passenger_queue")
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		PassengerMessageServer = mdw(PassengerMessageServer)
	}
	return nil
}

