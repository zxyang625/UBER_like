package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net/http"
	"pkg/dao/mq"
	"pkg/gateway"
	"strings"
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
	BillingMessageServer, err = mq.NewMessageServer("billing_queue")
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
	PayRespMessageServer, err = mq.NewMessageServer("")
	if err != nil {
		return err
	}
	for _, mdw := range mdws {
		PayRespMessageServer = mdw(PayRespMessageServer)
	}
	return nil
}

func ConsumePayQueue(ctx context.Context, logger log.Logger, tracer *zipkin.Tracer) {
	deliverServer := mq.InitDeliverMiddleware(tracer, "billing")(mq.HandleFunc(func(ctx context.Context, d amqp.Delivery) {
		d.Ack(false)
		asyncReq := mq.AsyncReq{}
		err := json.Unmarshal(d.Body, &asyncReq)
		if err != nil {
			logger.Log("mehtod", "json.Unmarshal", "err", err)
			return
		}

		url := strings.Join([]string{gateway.GatewayURL, asyncReq.Application, asyncReq.Service}, "/")
		httpReq, err := http.NewRequest(asyncReq.Method, url, bytes.NewBuffer(asyncReq.Data))
		httpReq.Header.Set("Content-type", "application/grpc")
		httpReq.URL.Scheme = "http"
		if err != nil {
			logger.Log("method", "NewRequest", "err", err)
			return
		}
		for k, v := range asyncReq.Header {
			httpReq.Header.Set(k, v)
		}
		httpReq.Header.Set("Length", fmt.Sprintf("%d", asyncReq.Priority))
		httpReq.Header.Set("Trace-ID", asyncReq.TraceID.String())
		rsp, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			logger.Log("method", "DefaultClient.Do", "err", err)
			return
		}
		rspData, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			logger.Log("method", "ioutil.ReadAll", "err", err)
			return
		}

		err = PayRespMessageServer.SendResp(ctx, d.ReplyTo, d.CorrelationId, rspData)
		if err != nil {
			logger.Log("method", "SendResp", "err", err)
			return
		}
	}))
	BillingMessageServer.Consume(ctx, ConsumePayQueueName, deliverServer)
}