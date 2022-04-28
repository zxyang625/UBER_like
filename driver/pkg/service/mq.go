package service

import (
	"pkg/dao/mq"
)

var DriverMessageServer mq.MessageServer

const (
	PublishQueueName = "driver_queue"
	ConsumeQueueName = "notification_queue"
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	DriverMessageServer, err = mq.NewMessageServer("driver_queue", 3)
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		DriverMessageServer = mdw(DriverMessageServer)
	}
	return nil
}
