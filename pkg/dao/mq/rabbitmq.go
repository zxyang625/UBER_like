package mq

import (
	"context"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
	Err "pkg/error"
	"strconv"
)

const (
	URI = "amqp://guest:guest@localhost:5672/"
)

type MessageServer interface {
	Publish(ctx context.Context, name string, priority int, data []byte) (err error)
	Consume(ctx context.Context, name string, handler DeliverHandler) error
	ReceiveResp(ctx context.Context) (d amqp.Delivery, err error)
	SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error)
}

func NewMessageServer(name string, length int) (MessageServer, error) {
	mq := &MQ{}

	var err error
	mq.Conn, err = amqp.Dial(URI)
	if err != nil {
		return nil, Err.New(Err.MQNewConnectionFail, err.Error())
	}

	for i := 0; i < length; i++ {
		ch, err := mq.Conn.Channel()
		if err != nil {
			return nil, Err.New(Err.MQInitChannelFail, err.Error())
		}
		mq.PublishChannels = append(mq.PublishChannels, ch)
	}
	mq.ConsumeChannel, err = mq.Conn.Channel()
	if err != nil {
		return nil, Err.New(Err.MQInitChannelFail, err.Error())
	}

	//mq.PublishChannel.Qos(
	//	//每次队列只消费一个消息 这个消息处理不完服务器不会发送第二个消息过来
	//	//当前消费者一次能接受的最大消息数量
	//	1,
	//	//服务器传递的最大容量
	//	0,
	//	//如果为true 对channel可用 false则只对当前队列可用
	//	true,
	//)
	//
	//mq.ConsumeChannel.Qos(
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

	mq.Q, err = mq.ConsumeChannel.QueueDeclare(
		name+"_reply",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, Err.New(Err.MQDeclareQueueFail, err.Error())
	}
	mq.Msgs, err = mq.ConsumeChannel.Consume(
		mq.Q.Name,
		"",
		false, //autoack
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

func (mq *MQ) Publish(ctx context.Context, name string, priority int, data []byte) (err error) {
	name = name + "_" + strconv.Itoa(priority)
	_, err = mq.PublishChannels[priority-1].QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return Err.New(Err.MQDeclareQueueFail, err.Error())
	}

	err = mq.PublishChannels[priority].Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: mq.CorrId,
			ReplyTo:       mq.Q.Name,
			Body:          data,
		},
	)
	if err != nil {
		return Err.New(Err.MQPublishMsgFail, err.Error())
	}
	return nil
}

func (mq *MQ) Consume(ctx context.Context, name string, handler DeliverHandler) error {
	q, err := mq.PublishChannels[0].QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return Err.New(Err.MQDeclareQueueFail, err.Error())
	}
	msgs, err := mq.PublishChannels[0].Consume(
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

	for d := range msgs {
		handler.Deliver(ctx, d)
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
	err = mq.PublishChannels[0].Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			Body:          data,
		})
	if err != nil {
		return Err.New(Err.MQSendRespFail, err.Error())
	}
	return nil
}
