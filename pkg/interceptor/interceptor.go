package interceptor

import (
	"context"
	"fmt"
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
	res := md.Get("Priority")
	if len(res) == 0 {
		return nil, Err.New(Err.InterceptorInvalidMeta, "priority value not found")
	}
	num, _ := strconv.Atoi(res[0])
	ctx1 := context.WithValue(ctx, "Priority", num + 1)
	fmt.Println(ctx1.Value("Priority"))
	// 继续处理请求
	return handler(ctx1, req)
}
