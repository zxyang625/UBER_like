package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"os"
	"pkg/dao/mq"
	"testing"
	"time"
)

const (
	Buf = 20.0
)

var testData = &mq.AsyncReq{
	Method:        "Method",
	OriginApp:     "OriginApp",
	OriginService: "OriginService",
	DestApp:       "DestApp",
	DestService:   "DestService",
	TraceID:       model.TraceID{},
	Length:        0,
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
	_, err = Ch.QueueDeclare("test_consumer_1", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("test_consumer_2", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("test_consumer_3", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	_, err = Ch.QueueDeclare("test_consumer_4", false, false, false, false, nil)
	if err != nil {
		t.Error(err)
	}
	data, _ := json.Marshal(testData)
	for i := 0; i < 50; i++ {
		Ch.Publish("", "test_consumer_1", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	}
	for i := 0; i < 100; i++ {
		Ch.Publish("", "test_consumer_2", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	}
	for i := 0; i < 100; i++ {
		Ch.Publish("", "test_consumer_3", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	}
	for i := 0; i < 50; i++ {
		Ch.Publish("", "test_consumer_4", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	}
	Ch.Close()
}

var start time.Time
var buf bytes.Buffer
var str string

func TestConsumer(t *testing.T) {
	consumer, err := InitQueueServer(4, "test_consumer")
	if err != nil {
		t.Error(err)
	}
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		t.Error(err)
	}
	start = time.Now()
	consumer.TestPPolling()
	time.Sleep(20 * time.Second)
	f.WriteString(buf.String())
}

// Priority polling based on multi queue
func (c *ConsumeQueueServer) TestPPolling() {
	list := [5]int{0, 0, 0, 0, 0}
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				for i := 1; i <= c.ConsumeQueueSize; i++ {
					m := float32(len(c.DeliveryMap[i]))
					c.qavg = (1.0-c.weights[i-1])*c.qavg + c.weights[i-1]*m
					var max int
					switch {
					case Buf-m < m-c.qavg:
						max = i + 1
					case m == 0:
						max = 0
					default:
						max = 1
					}

					//max := 1
					//if m == 0 {
					//	max = 0
					//}
					list[i] = max
					for j := 0; j < max; j++ {
						d := <-c.DeliveryMap[i]
						d.Ack(false)
					}
					time.Sleep(40 * time.Millisecond)
				}
				str = fmt.Sprintf("%v %v %v %v %v %v\n", time.Since(start).String(), list[1], list[2], list[3], list[4], c.qavg)
				buf.WriteString(str)
			}
		}()
	}
}
