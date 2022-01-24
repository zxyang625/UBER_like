package mq

import (
	"context"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
)

type MQ struct {
	Conn             *amqp.Connection
	Ch               *amqp.Channel
	Q                amqp.Queue
	Msgs             <-chan amqp.Delivery
	CorrId           string
}

type MQModel struct {
	Data      []byte          `json:"data"`
	SpanModel model.SpanModel `json:"span_model"`
}

type SpanModel struct {
	TraceID  string `json:"trace_id"`
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

type DeliverHandler interface {
	Deliver(context.Context, amqp.Delivery)
}

type HandleFunc func(ctx context.Context, d amqp.Delivery)
func (h HandleFunc) Deliver(ctx context.Context, d amqp.Delivery) {
	h(ctx, d)
}