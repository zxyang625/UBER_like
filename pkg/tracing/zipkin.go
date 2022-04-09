package tracing

import (
	"github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
	Err "pkg/error"
)

const DefaultZipkinURL = "http://localhost:9411/api/v2/spans"

type TracingImpl struct {
	NativeTracer *zipkin.Tracer
	Reporter reporter.Reporter
	opentracing.Tracer
}

func NewOpenTracingTracer(serviceName string) (*TracingImpl, error){
	if serviceName == "" {
		return nil, Err.New(Err.TracingEmptyService, "empty service name")
	}
	tracer := new(TracingImpl)
	tracer.Reporter = http.NewReporter(DefaultZipkinURL)
	endpoint, err := zipkin.NewEndpoint(serviceName, "")
	if err != nil {
		return nil, Err.Errorf(Err.TracingNewTracerFail, "new endpoint for service: %s failed", serviceName)
	}
	nativeTracer, err := zipkin.NewTracer(tracer.Reporter, zipkin.WithLocalEndpoint(endpoint), zipkin.WithSharedSpans(true))
	if err != nil {
		return nil, Err.Errorf(Err.TracingNewTracerFail, "NewTracer failed, err: %v", err)
	}
	tracer.NativeTracer = nativeTracer
	tracer.Tracer = zipkintracer.Wrap(nativeTracer)
	return tracer, nil
}
