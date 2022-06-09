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
	"pkg/util"
	"strconv"
	"strings"
	"time"
)

const (
	RabbitMQURI = "amqp://guest:guest@localhost:5672/"
	defaultQavg = 0
	defaultBuf  = 50
)

type ConsumeQueueServer struct {
	Conn             *amqp.Connection
	ChMap            map[int]*amqp.Channel
	PublishChannel   *amqp.Channel
	DeliveryMap      map[int]<-chan amqp.Delivery
	ConsumeQueueSize int
	ConsumeQueueName string
	weights          []float32
	qavg             float32
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
	loadbalancer := loadbalance.NewLoadBalancer()
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
		instance, err := loadbalancer.RandomSelect(instances)
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", instance.Service.Address, instance.Service.Port)
		req.URL.Path = "/" + destPath

		length := req.Header.Get("Length")
		if length == "" {
			req.Header.Set("Length", "0")
		} else {
			num, _ := strconv.Atoi(length)
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
	server.weights = make([]float32, consumeQueueSize)
	var total float32
	for i := 1; i <= consumeQueueSize; i++ {
		total += float32(i)
		server.ChMap[i], err = server.Conn.Channel()
		if err != nil {
			return nil, Err.New(Err.MQInitChannelFail, err.Error())
		}
		server.DeliveryMap[i], err = server.ConsumeSingleQueue(consumeQueueName, i)
		if err != nil {
			return nil, Err.New(Err.MQConsumeMsgFail, err.Error())
		}
	}
	server.PublishChannel, _ = server.Conn.Channel()
	for i := 1; i <= consumeQueueSize; i++ {
		server.weights[i-1] = float32(i) / total
	}
	server.qavg = defaultQavg
	return server, nil
}

func (c *ConsumeQueueServer) ConsumeSingleQueue(queueName string, queuePriority int) (<-chan amqp.Delivery, error) {
	name := strings.Join([]string{queueName, strconv.Itoa(queuePriority)}, "_")
	_, err := c.ChMap[queuePriority].QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		log.Printf("declare queue failed, name: %s, err: %v", name, err)
		return nil, err
	}
	d, err := c.ChMap[queuePriority].Consume(name, "", false, false, false, false, nil)
	if err != nil {
		log.Printf("consume queue failed, name: %s, err: %v", name, err)
		return nil, err
	}
	log.Printf("consuming queue: %s, priority: %d", queueName, queuePriority)
	return d, nil
}

func (c *ConsumeQueueServer) SendResp(queueServerData *QueueServerData, rspData []byte) error {
	//err := s.ChMap[queueServerData.Req.Length].Publish(
	err := c.PublishChannel.Publish(
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

func (c *ConsumeQueueServer) Consume(serviceName string, size int) {
	for worker := 0; worker < 4; worker++ {
		go func() {
			list := make([]mq.AsyncReq, 0, 5)
			for {
				for i := 1; i <= c.ConsumeQueueSize; i++ {
					m := float32(len(c.DeliveryMap[i]))
					c.qavg = (1.0-c.weights[i-1])*c.qavg + c.weights[i-1]*m
					var max int
					switch {
					case defaultBuf-m < m-c.qavg:
						max = i + 1
					case m == 0:
						max = 0
						time.Sleep(2 * time.Millisecond)
					default:
						max = 1
					}

					req := mq.AsyncReq{}
					var replyTo, corrId string
					for j := 0; j < max; j++ {
						d := <-c.DeliveryMap[i]
						d.Ack(false)
						replyTo, corrId = d.ReplyTo, d.CorrelationId
						if d.Body == nil {
							fmt.Println("error empty Body")
							continue
						}
						err := json.Unmarshal(d.Body, &req)
						if err != nil {
							fmt.Printf("json unmarshal failed, service: %s, req: %+v, err: %v", serviceName, req, err)
							continue
						}
						list = append(list, req)
					}
					for _, v := range list {
						ReqBufferChan <- QueueServerData{
							ReplyTo: replyTo,
							CorrId:  corrId,
							Req:     v,
						}
					}
					list = list[0:0]
				}
			}

			//for {
			//	for i := 1; i <= c.ConsumeQueueSize; i++ {
			//		m := float32(len(c.DeliveryMap[i]))
			//		if m == 0.0 {
			//			time.Sleep(3 * time.Millisecond)
			//			continue
			//		}
			//		req := mq.AsyncReq{}
			//		d := <-c.DeliveryMap[i]
			//		d.Ack(false)
			//		if d.Body == nil {
			//			fmt.Println("error empty Body")
			//			continue
			//		}
			//		err := json.Unmarshal(d.Body, &req)
			//		if err != nil {
			//			fmt.Printf("json unmarshal failed, service: %s, req: %+v, err: %v", serviceName, req, err)
			//			continue
			//		}
			//		ReqBufferChan <- QueueServerData{
			//			ReplyTo: d.ReplyTo,
			//			CorrId:  d.CorrelationId,
			//			Req:     req,
			//		}
			//	}
			//}
		}()
	}
}

func ProxySendReq(logger kitlog.Logger, s *ConsumeQueueServer, tracer *zipkin.Tracer, gatewayUrl string) {
	for i := 0; i < 4; i++ {
		go func() {
			for {
				queueServerData := <-ReqBufferChan
				Req := queueServerData.Req

				traceID := Req.TraceID
				priority := util.CalcPriority(Req.OriginApp, Req.DestApp, Req.Length)
				span := tracer.StartSpan("gateway", zipkin.Parent(model.SpanContext{TraceID: traceID}))
				span.Tag("Length", strconv.Itoa(Req.Length))
				span.Tag("Priority", strconv.Itoa(priority))

				url := strings.Join([]string{gatewayUrl, Req.DestApp, Req.DestService}, "/")
				if Req.DestService == "get-bill" {
					url = url + "/250205"
				}
				httpReq, err := http.NewRequest(Req.Method, url, bytes.NewBuffer(Req.Data))
				httpReq.Header.Set("Content-type", "application/grpc")
				httpReq.URL.Scheme = "http"
				if err != nil {
					fmt.Println("method", "NewRequest", "err", err)
					span.Finish()
					continue
				}
				for k, v := range Req.Header {
					httpReq.Header.Set(k, v)
				}
				httpReq.Header.Set("Length", fmt.Sprintf("%d", Req.Length))
				httpReq.Header.Set("Priority", fmt.Sprintf("%d", priority))
				httpReq.Header.Set("Trace-ID", Req.TraceID.String())
				rsp, err := http.DefaultClient.Do(httpReq)
				if err != nil {
					fmt.Println("method", "DefaultClient.Do", "err", err)
					span.Finish()
					continue
				}
				_, err = ioutil.ReadAll(rsp.Body)
				if err != nil {
					fmt.Println("method", "ioutil.ReadAll", "err", err)
					span.Finish()
					continue
				}
				span.Finish()

				//err = s.SendResp(&queueServerData, rspData)
				//if err != nil {
				//	fmt.Println("send resp", "failed,", "replyto", queueServerData.ReplyTo, "corrId", queueServerData.CorrId)
				//	continue
				//}
			}
		}()
	}
}
