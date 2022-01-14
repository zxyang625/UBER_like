package endpoint

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	"payment/pkg/config"
	"time"

	endpoint "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With(config.System+"_histogram", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func CountingMiddleware(count metrics.Counter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				count.With(config.System+"_counter", fmt.Sprint(err == nil)).Add(1)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin).Microseconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func TracingMiddle(methodName string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			parentSpan := opentracing.SpanFromContext(ctx)
			childSpan := parentSpan.Tracer().StartSpan("service." + methodName, opentracing.ChildOf(parentSpan.Context()))
			defer childSpan.Finish()
			return next(ctx, request)
		}
	}
}


func TraceEndpoint(tracer *zipkin.Tracer, name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			var sc model.SpanContext
			if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
				sc = parentSpan.Context()
			}
			sp := tracer.StartSpan(name, zipkin.Parent(sc))
			defer sp.Finish()

			ctx = zipkin.NewContext(ctx, sp)
			return next(ctx, request)
		}
	}
}