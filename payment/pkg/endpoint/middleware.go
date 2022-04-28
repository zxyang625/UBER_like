package endpoint

import (
	"context"
	"fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	metrics "github.com/go-kit/kit/metrics"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"github.com/openzipkin/zipkin-go/model"
	"pkg/config"
	"strconv"
	"time"
)

// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				duration.With(config.SystemPayment+"_histogram", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func CountingMiddleware(count metrics.Counter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				count.With(config.SystemPayment+"_counter", fmt.Sprint(err == nil)).Add(1)
			}(time.Now())
			return next(ctx, request)
		}
	}
}

// LoggingMiddleware returns an endpoint middleware that logs the
// duration of each invocation, and the resulting error, if any.
func LoggingMiddleware(logger kitlog.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin).Microseconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func TraceEndpoint(tracer *zipkin.Tracer, name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			//var sc model.SpanContext
			//traceStr := ctx.Value("Trace-ID")
			//if traceStr != nil {
			//	traceID, _ := model.TraceIDFromHex(traceStr.(string))
			//	sc.TraceID = traceID
			//} else {
			//	if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
			//		sc = parentSpan.Context()
			//	}
			//}
			//span := tracer.StartSpan(name, zipkin.Parent(sc))
			//defer span.Finish()
			traceStr := ctx.Value("Trace-ID")
			traceID := model.TraceID{}
			if traceStr == nil {
				traceID = idgenerator.NewRandom64().TraceID()
			} else {
				traceID, _ = model.TraceIDFromHex(traceStr.(string))
			}
			span := tracer.StartSpan(name, zipkin.Parent(model.SpanContext{TraceID: traceID}))
			span.Tag("Length", strconv.Itoa(ctx.Value("Length").(int)))
			defer span.Finish()
			return next(ctx, request)
		}
	}
}
