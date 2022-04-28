package gateway

import (
	"encoding/json"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"pkg/dao/mq"
	"testing"
)

var testData = &mq.AsyncReq{
	Method:        "Method",
	OriginApp:     "OriginApp",
	OriginService: "OriginService",
	DestApp:       "DestApp",
	DestService:   "DestService",
	TraceID:       model.TraceID{},
	Priority:      0,
	Header:        nil,
	Data:          []byte("hello world"),
}

func TestProvider(t *testing.T) {
	Conn, err := amqp.Dial(RabbitMQURI)
	if err != nil {
		t.Error(err)
	}
	Ch, err := Conn.Channel()
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("pay_queue_1", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("pay_queue_2", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("pay_queue_3", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	data, _ := json.Marshal(testData)
	for i := 0; i < 20; i++ {
		Ch.Publish("", "pay_queue_1", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
		Ch.Publish("", "pay_queue_2", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
		Ch.Publish("", "pay_queue_3", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	}
	Ch.Close()
}
