package config

import (
	"application/common/xes"
	"application/common/xzap"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	DetailsKqConsumer kq.KqConf
	BindsKqConsumer   kq.KqConf
	ES                xes.Config
	Zap               xzap.Config
}
