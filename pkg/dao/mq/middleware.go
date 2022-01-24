package mq

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
)

type Middleware func(MessageServer) MessageServer
type DeliverMiddleware func(DeliverHandler) DeliverHandler

type loggingMiddleware struct {
	next MessageServer
	logger log.Logger
}

type tracingMiddleware struct {
	next MessageServer
	name string
	tracer *zipkin.Tracer
}

type DeliveringMiddleware struct {
	next DeliverHandler
	name string
	tracer *zipkin.Tracer
}

func (d DeliveringMiddleware) Deliver(ctx context.Context, delivery amqp.Delivery) {
	mqModel := &MQModel{}
	err := json.Unmarshal(delivery.Body, mqModel)
	if err != nil {
		return
	}
	span := d.tracer.StartSpan(d.name + "/consume", zipkin.Parent(model.SpanContext{
		TraceID:  mqModel.SpanModel.TraceID,
		//ID:       mqModel.SpanModel.ID,
		//ParentID: mqModel.SpanModel.ParentID,
	}))
	defer span.Finish()
	ctx = zipkin.NewContext(ctx, span)
	d.next.Deliver(ctx, delivery)
}

func InitLoggingMiddleware(logger log.Logger) Middleware {
	return func(next MessageServer) MessageServer {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func InitTracingMiddleware(tracer *zipkin.Tracer, name string) Middleware {
	return func(next MessageServer) MessageServer {
		return &tracingMiddleware{
			next: next,
			name: name,
			tracer: tracer,
		}
	}
}

func InitDeliverMiddleware(tracer *zipkin.Tracer, name string) DeliverMiddleware {
	return func(next DeliverHandler) DeliverHandler {
		return &DeliveringMiddleware{
			next: next,
			name: name,
			tracer: tracer,
		}
	}
}

func (t *tracingMiddleware) Publish(ctx context.Context, name string, data []byte) (err error) {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "publish", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.Publish(ctx, name, data)
}

func (t *tracingMiddleware) Consume(ctx context.Context, name string, handler DeliverHandler) error {
	//var sc model.SpanContext
	//if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
	//	sc = parentSpan.Context()
	//}
	//sp := t.tracer.StartSpan(t.name + "consume", zipkin.Parent(sc))
	//defer sp.Finish()
	//
	//ctx = zipkin.NewContext(ctx, sp)
	return t.next.Consume(ctx, name, handler)
}

func (t *tracingMiddleware) ReceiveResp(ctx context.Context) (d amqp.Delivery, err error) {
	//var sc model.SpanContext
	//if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
	//	sc = parentSpan.Context()
	//}
	//sp := t.tracer.StartSpan(t.name + "receive_resp", zipkin.Parent(sc))
	//defer sp.Finish()
	//
	//ctx = zipkin.NewContext(ctx, sp)
	return t.next.ReceiveResp(ctx)
}

func (t *tracingMiddleware) SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error) {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "send_resp", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.SendResp(ctx, routingKey, corrId, data)
}

func (m *loggingMiddleware) Publish(ctx context.Context, name string, data []byte) (err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "publish", "name", name, "err", err)
	}()
	return m.next.Publish(ctx, name, data)
}

func (m *loggingMiddleware) Consume(ctx context.Context, name string, handler DeliverHandler) (err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "Consume", "name", name, "err", err)
	}()
	return m.next.Consume(ctx, name, handler)
}


func (m *loggingMiddleware) ReceiveResp(ctx context.Context) (d amqp.Delivery, err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "ReceiveResp", "err", err)
	}()
	return m.next.ReceiveResp(ctx)
}

func (m *loggingMiddleware) SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "SendResp", "err", err)
	}()
	return m.next.SendResp(ctx, routingKey, corrId, data)
}
