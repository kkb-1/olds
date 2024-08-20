package svc

import (
	"application/apps/user-manage/mq/internal/config"
	"application/common/xes"
	"application/common/xzap"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

type ServiceContext struct {
	Config   config.Config
	ESClient *elasticsearch.Client
	Logger   *zap.Logger
}

func NewServiceContext(c config.Config) *ServiceContext {
	esClient := xes.MustNew(c.ES)
	logger := xzap.New(c.Zap)
	return &ServiceContext{
		Config:   c,
		ESClient: esClient,
		Logger:   logger,
	}
}
