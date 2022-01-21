package mq

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"github.com/streadway/amqp"
)

type Middleware func(MessageServer) MessageServer

type loggingMiddleware struct {
	next MessageServer
	logger log.Logger
}

type tracingMiddleware struct {
	next MessageServer
	name string
	tracer *zipkin.Tracer
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

func (m *loggingMiddleware) Publish(ctx context.Context, name string, data []byte) (err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "publish", "err", err)
	}()
	return m.next.Publish(ctx, name, data)
}

func (m *loggingMiddleware) Consume(ctx context.Context, name string, handlerFunc func(d amqp.Delivery)) (err error) {
	defer func() {
		m.logger.Log("system", "RabbitMQ", "method", "Consume", "err", err)
	}()
	return m.next.Consume(ctx, name, handlerFunc)
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

func (t *tracingMiddleware) Publish(ctx context.Context, name string, data []byte) (err error) {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "_publish", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.Publish(ctx, name, data)
}

func (t *tracingMiddleware) Consume(ctx context.Context, name string, handlerFunc func(d amqp.Delivery)) error {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "_consume", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.Consume(ctx, name, handlerFunc)
}

func (t *tracingMiddleware) ReceiveResp(ctx context.Context) (d amqp.Delivery, err error) {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "_receive_resp", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.ReceiveResp(ctx)
}

func (t *tracingMiddleware) SendResp(ctx context.Context, routingKey, corrId string, data []byte) (err error) {
	var sc model.SpanContext
	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
		sc = parentSpan.Context()
	}
	sp := t.tracer.StartSpan(t.name + "_send_resp", zipkin.Parent(sc))
	defer sp.Finish()

	ctx = zipkin.NewContext(ctx, sp)
	return t.next.SendResp(ctx, routingKey, corrId, data)
}
