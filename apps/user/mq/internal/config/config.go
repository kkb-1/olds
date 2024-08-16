package config

import (
	"application/common/xes"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	KqConsumerConf kq.KqConf
	ES             xes.Config
}
