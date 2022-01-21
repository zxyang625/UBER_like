package mq

import (
	"context"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	Err "pkg/error"
)

const (
	URI                 = "amqp://guest:guest@localhost:5672/"
	PassengerQueueName = "passenger"
	TripQueueName = "trip"
	TripCorrId    = "passenger_to_trip_corrID"
)

type MessageServer interface {
	Publish(ctx context.Context, name string, data []byte) (err error)
	Consume(ctx context.Context, name string, handlerFunc func(d amqp.Delivery)) error
	ReceiveResp(ctx context.Context) (d amqp.Delivery, err error)
	SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error)
}

type MQ struct {
	Conn             *amqp.Connection
	Ch               *amqp.Channel
	Q                amqp.Queue
	Msgs             <-chan amqp.Delivery
	CorrId           string
}

func NewMessageServer(name string) (MessageServer, error) {
	mq := &MQ{}

	var err error
	mq.Conn, err = amqp.Dial(URI)
	if err != nil {
		return nil, Err.New(Err.MQNewConnectionFail, err.Error())
	}

	mq.Ch, err = mq.Conn.Channel()
	if err != nil {
		return nil, Err.New(Err.MQInitChannelFail, err.Error())
	}

	//mq.Ch.Qos(
	//	//每次队列只消费一个消息 这个消息处理不完服务器不会发送第二个消息过来
	//	//当前消费者一次能接受的最大消息数量
	//	1,
	//	//服务器传递的最大容量
	//	0,
	//	//如果为true 对channel可用 false则只对当前队列可用
	//	true,
	//)

	if name == "" {
		log.Println("init resp sender success")
		return mq, nil
	}

	mq.Q, err = mq.Ch.QueueDeclare(
		name + "_reply",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, Err.New(Err.MQDeclareQueueFail, err.Error())
	}
	mq.Msgs, err = mq.Ch.Consume(
		mq.Q.Name,
		"",
		false,			//autoack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, Err.New(Err.MQNewMsgsFail, err.Error())
	}

	mq.CorrId = name + uuid.New().String()

	log.Printf("init MQ success, name: %s\n", mq.Q.Name)
	return mq, nil
}

func (mq *MQ) Publish(ctx context.Context, name string, data []byte) (err error) {
	_, err = mq.Ch.QueueDeclare(
		name,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return Err.New(Err.MQDeclareQueueFail, err.Error())
	}

	err = mq.Ch.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			CorrelationId:   mq.CorrId,
			ReplyTo:         mq.Q.Name,
			Body:            data,
		},
	)
	if err != nil {
		return Err.New(Err.MQPublishMsgFail, err.Error())
	}
	return nil
}

func (mq *MQ) Consume(ctx context.Context, name string, handlerFunc func(d amqp.Delivery)) error {
	q, err := mq.Ch.QueueDeclare(
		name,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return Err.New(Err.MQDeclareQueueFail, err.Error())
	}
	msgs, err := mq.Ch.Consume(
		q.Name,
	"",
	false,
	false,
	false,
	false,
	nil,
	)
	if err != nil {
		return Err.New(Err.MQConsumeMsgFail, err.Error())
	}

	for d := range msgs{
		handlerFunc(d)
	}
	return nil
}

func (mq *MQ) ReceiveResp(ctx context.Context) (d amqp.Delivery, err error) {
	d = <-mq.Msgs
	defer d.Ack(false)
	//if !ok {
	//	return amqp.Delivery{}, Err.New(Err.MQGetRespFail, "ReceiveResp failed")
	//}
	return d, nil
}


func (mq *MQ) SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error) {
	err = mq.Ch.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			CorrelationId: corrId,
			Body: data,
		})
	if err != nil {
		return Err.New(Err.MQSendRespFail, err.Error())
	}
	return nil
}