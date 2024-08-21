package config

import (
	"application/common/xzap"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRPC       zrpc.RpcClientConf
	UserManageRPC zrpc.RpcClientConf
	Zap           xzap.Config
}
