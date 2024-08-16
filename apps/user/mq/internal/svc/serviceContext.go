package svc

import (
	"application/apps/user/mq/internal/config"
	"application/common/xes"
	"github.com/elastic/go-elasticsearch/v8"
)

type ServiceContext struct {
	Config   config.Config
	ESClient *elasticsearch.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	esClient := xes.MustNew(c.ES)
	return &ServiceContext{
		Config:   c,
		ESClient: esClient,
	}
}
