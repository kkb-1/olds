package interceptors

import (
	"context"

	"application/common/xcode"

	"google.golang.org/grpc"
)

/*
rpc server拦截器
使用说明：
请看app/resp_api/rpc/user.go

	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())
	注册一下拦截器就OK了
*/
func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.FromError(err).Err()
	}
}
