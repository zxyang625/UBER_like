package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	kitlog "github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"pkg/dao/mq"
	"pkg/discover"
	Err "pkg/error"
	"pkg/loadbalance"
	"strconv"
	"strings"
	"time"
)

const (
	RabbitMQURI = "amqp://guest:guest@localhost:5672/"
)

type ConsumeQueueServer struct {
	Conn             *amqp.Connection
	ChMap            map[int]*amqp.Channel
	PublishChannel   *amqp.Channel
	DeliveryMap      map[int]<-chan amqp.Delivery
	ConsumeQueueSize int
	ConsumeQueueName string
}

type QueueServerData struct {
	ReplyTo string
	CorrId  string
	Req     mq.AsyncReq
}

var ReqBufferChan = make(chan QueueServerData)

func NewReverseProxy(consulHost string, consulPort int, logger kitlog.Logger) (*httputil.ReverseProxy, error) {
	client, err := discover.NewDiscoverClient(consulHost, consulPort, true)
	if err != nil {
		logger.Log("NewDiscoverClient", "fail", "err", err)
		return nil, err
	}
	director := func(req *http.Request) {
		reqPath := req.URL.Path
		if reqPath == "" {
			logger.Log("method", "NewReverseProxy", "err", "empty url req path")
			return
		}
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]
		instances, err := client.DiscoverServices(serviceName, "", true)
		if err != nil {
			logger.Log("service name", serviceName, "msg", "query instances failed", "err", err)
			return
		}

		if len(instances) == 0 {
			logger.Log("service name", serviceName, "err", "no such service instance")
			return
		}

		destPath := strings.Join(pathArray[1:], "/")
		addr, port, err := loadbalance.NewLoadBalancer().Select(instances)
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", addr, port)
		req.URL.Path = "/" + destPath

		priority := req.Header.Get("Length")
		if priority == "" {
			req.Header.Set("Length", "0")
		} else {
			num, _ := strconv.Atoi(priority)
			req.Header.Set("Length", fmt.Sprintf("%d", num))
		}
	}
	return &httputil.ReverseProxy{
		Director: director,
	}, nil
}

func InitQueueServer(consumeQueueSize int, consumeQueueName string) (*ConsumeQueueServer, error) {
	server := &ConsumeQueueServer{ConsumeQueueSize: consumeQueueSize, ChMap: map[int]*amqp.Channel{}, DeliveryMap: map[int]<-chan amqp.Delivery{}}
	var err error
	server.Conn, err = amqp.Dial(RabbitMQURI)
	if err != nil {
		return nil, Err.New(Err.MQNewConnectionFail, err.Error())
	}
	for i := 1; i <= consumeQueueSize; i++ {
		server.ChMap[i], err = server.Conn.Channel()
		if err != nil {
			return nil, Err.New(Err.MQInitChannelFail, err.Error())
		}
		server.DeliveryMap[i], err = server.ConsumeSingleQueue(consumeQueueName, i)
	}
	server.PublishChannel, _ = server.Conn.Channel()
	return server, nil
}

func (s *ConsumeQueueServer) ConsumeSingleQueue(queueName string, queuePriority int) (<-chan amqp.Delivery, error) {
	name := strings.Join([]string{queueName, strconv.Itoa(queuePriority)}, "_")
	_, err := s.ChMap[queuePriority].QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		log.Printf("declare queue failed, name: %s, err: %v", name, err)
		return nil, err
	}
	d, err := s.ChMap[queuePriority].Consume(name, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("consume queue failed, name: %s, err: %v", name, err)
		return nil, err
	}
	log.Printf("consuming queue: %s, priority: %d", queueName, queuePriority)
	return d, nil
}

func (s *ConsumeQueueServer) SendResp(queueServerData *QueueServerData, rspData []byte) error {
	//err := s.ChMap[queueServerData.Req.Priority].Publish(
	err := s.PublishChannel.Publish(
		"",
		queueServerData.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: queueServerData.CorrId,
			Body:          rspData,
		})
	if err != nil {
		return err
	}
	return nil
}

func (s *ConsumeQueueServer) Consume(serviceName string) {
	for t := 0; t < 3; t++ {
		tick := time.Tick(time.Millisecond * 5)
		go func() {
			i := 0
			for {
				i = i%s.ConsumeQueueSize + 1
				select {
				case d := <-s.DeliveryMap[i]:
					req := mq.AsyncReq{}
					if d.Body == nil {
						d.Ack(false)
						break
					}
					err := json.Unmarshal(d.Body, &req)
					if err != nil {
						fmt.Printf("json unmarshal failed, service: %s, req: %+v, err: %v", serviceName, req, err)
						d.Ack(false)
						break
					}
					d.Ack(false)
					ReqBufferChan <- QueueServerData{
						ReplyTo: d.ReplyTo,
						CorrId:  d.CorrelationId,
						Req:     req,
					}
					for d := range s.DeliveryMap[i] {
						req := mq.AsyncReq{}
						if d.Body == nil {
							d.Ack(false)
							break
						}
						err := json.Unmarshal(d.Body, &req)
						if err != nil {
							fmt.Printf("json unmarshal failed, service: %s, req: %+v, err: %v", serviceName, req, err)
							d.Ack(false)
							break
						}
						d.Ack(false)
						ReqBufferChan <- QueueServerData{
							ReplyTo: d.ReplyTo,
							CorrId:  d.CorrelationId,
							Req:     req,
						}
					}
				case <-tick:
					break
				}
			}
		}()
	}
}

func ProxySendReq(logger kitlog.Logger, s *ConsumeQueueServer, tracer *zipkin.Tracer, gatewayUrl string) {
	for i := 0; i < 3; i++ {
		go func() {
			for {
				queueServerData := <-ReqBufferChan
				Req := queueServerData.Req
				traceID := Req.TraceID
				span := tracer.StartSpan("gateway", zipkin.Parent(model.SpanContext{TraceID: traceID}))
				span.Tag("Length", strconv.Itoa(Req.Priority))

				url := strings.Join([]string{gatewayUrl, Req.DestApp, Req.DestService}, "/")
				httpReq, err := http.NewRequest(Req.Method, url, bytes.NewBuffer(Req.Data))
				httpReq.Header.Set("Content-type", "application/grpc")
				httpReq.URL.Scheme = "http"
				if err != nil {
					logger.Log("method", "NewRequest", "err", err)
					span.Finish()
					continue
				}
				for k, v := range Req.Header {
					httpReq.Header.Set(k, v)
				}
				httpReq.Header.Set("Length", fmt.Sprintf("%d", Req.Priority))
				httpReq.Header.Set("Trace-ID", Req.TraceID.String())
				rsp, err := http.DefaultClient.Do(httpReq)
				if err != nil {
					logger.Log("method", "DefaultClient.Do", "err", err)
					span.Finish()
					continue
				}
				rspData, err := ioutil.ReadAll(rsp.Body)
				if err != nil {
					logger.Log("method", "ioutil.ReadAll", "err", err)
					span.Finish()
					continue
				}
				span.Finish()

				err = s.SendResp(&queueServerData, rspData)
				if err != nil {
					logger.Log("send resp", "failed,", "replyto", queueServerData.ReplyTo, "corrId", queueServerData.CorrId)
					continue
				}
			}
		}()
	}
}
