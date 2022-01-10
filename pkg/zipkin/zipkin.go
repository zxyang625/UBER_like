package zipkin

import (
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	"github.com/openzipkin/zipkin-go/reporter/http"
	Err "pkg/error"
)

const DefaultZipkinURL = "http://localhost:9411/api/v2/spans"

func NewZipkinOption(reporter reporter.Reporter) (tracerOption grpc.ServerOption, err error){
	if reporter == nil {
		reporter = http.NewReporter(DefaultZipkinURL)
	}
	zkTracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		return nil, Err.Errorf(Err.ZipkinNewTracerFail, "NewZipkinReport failed, err: %v", err)
	}
	return kitzipkin.GRPCServerTrace(zkTracer), nil
}