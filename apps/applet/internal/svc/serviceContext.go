package svc

import (
	"application/apps/applet/internal/config"
	"application/apps/user-manage/rpc/userManageClient"
	"application/apps/user/rpc/userClient"
	"application/common/interceptors"
	"application/common/xzap"
	"github.com/zeromicro/go-zero/zrpc"
	"go.uber.org/zap"
)

type ServiceContext struct {
	Config        config.Config
	UserRPC       userClient.User
	UserManageRPC userManageClient.UserManage
	Logger        *zap.Logger
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := xzap.New(c.Zap, 1)
	user := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	userManage := zrpc.MustNewClient(c.UserManageRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:        c,
		UserRPC:       userClient.NewUser(user),
		UserManageRPC: userManageClient.NewUserManage(userManage),
		Logger:        logger,
	}
}
