package interceptor

import (
	"context"
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
	priority := md.Get("Length")
	if len(priority) == 0 {
		return nil, Err.New(Err.InterceptorInvalidMeta, "priority value not found")
	}
	num, _ := strconv.Atoi(priority[0])
	ctx = context.WithValue(ctx, "Length", num + 1)
	if len(md.Get("Trace-ID")) > 0 {
		ctx = context.WithValue(ctx, "Trace-ID", md.Get("Trace-ID")[0])
	}
	//继续处理请求
	return handler(ctx, req)
}
