package mq

import (
	"context"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
	"sync"
)

type MQ struct {
	Conn            *amqp.Connection
	mutex           sync.Mutex
	PublishChannels []*amqp.Channel
	ConsumeChannel  *amqp.Channel
	Q               amqp.Queue
	Msgs            <-chan amqp.Delivery
	CorrId          string
}

type MQModel struct {
	Data      []byte          `json:"data"`
	SpanModel model.SpanModel `json:"span_model"`
}

type DeliverHandler interface {
	Deliver(context.Context, amqp.Delivery)
}

type HandleFunc func(ctx context.Context, d amqp.Delivery)

func (h HandleFunc) Deliver(ctx context.Context, d amqp.Delivery) {
	h(ctx, d)
}

type AsyncReq struct {
	Method        string `json:"method"`
	OriginApp     string `json:"origin_app"`
	OriginService string `json:"origin_service"`
	DestApp       string `json:"dest_app"`
	DestService   string `json:"dest_service"`
	// URL         string            `json:"url"`
	TraceID  model.TraceID     `json:"trace_id"`
	Priority int               `json:"priority"`
	Header   map[string]string `json:"header"`
	Data     []byte            `json:"data"`
}
