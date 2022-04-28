package service

import (
	"pkg/dao/mq"
)

var PayMessageServer mq.MessageServer

const (
	PublishQueueName = "pay_queue"
	ConsumeQueueName = ""
)

func InitMessageServer(mdws ...mq.Middleware) error {
	var err error
	PayMessageServer, err = mq.NewMessageServer("pay_queue", 3)
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		PayMessageServer = mdw(PayMessageServer)
	}
	return nil
}
