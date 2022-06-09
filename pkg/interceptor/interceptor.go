package interceptor

import (
	"context"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	Err "pkg/error"
	"strconv"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/grpc.health.v1.Health/Check" {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, Err.New(Err.InterceptorInvalidCtx, "no metadata in context")
	}
	length := md.Get("Length")
	if len(length) == 0 {
		return nil, Err.New(Err.InterceptorInvalidMeta, "length value not found")
	}
	priority := md.Get("Priority")
	if len(priority) == 0 {
		priority = []string{"1"}
	}
	num, _ := strconv.Atoi(length[0])
	ctx = context.WithValue(ctx, "Length", num+1)
	p, _ := strconv.Atoi(priority[0])
	ctx = context.WithValue(ctx, "Priority", p)
	if md.Get("Trace-ID")[0] != "" {
		ctx = context.WithValue(ctx, "Trace-ID", md.Get("Trace-ID")[0])
	} else {
		ctx = context.WithValue(ctx, "Trace-ID", idgenerator.NewRandom64().TraceID().String())
	}
	//继续处理请求
	return handler(ctx, req)
}
