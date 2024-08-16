package config

import (
	"application/common/xgorm"
	"application/common/xredis"
	"application/common/xzap"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Zap    xzap.Config
	DB     xgorm.Mysql
	XRedis xredis.Config
	//ES    xes.Config
}
