package config

import (
	"application/common/xes"
	"application/common/xgorm"
	"application/common/xzap"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Zap xzap.Config
	DB  xgorm.Mysql
	ES  xes.Config
}
