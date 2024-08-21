package svc

import (
	"application/apps/user-manage/rpc/internal/config"
	"application/common/xes"
	"application/common/xgorm"
	"application/common/xzap"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	ES     *elasticsearch.TypedClient
	Logger *zap.Logger
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := xzap.New(c.Zap, 0)
	db := xgorm.MustOpen(c.DB, xzap.GetGormLog(xzap.New(c.Zap, 2)))
	es := xes.MustNewType(c.ES)
	return &ServiceContext{
		Config: c,
		DB:     db,
		ES:     es,
		Logger: logger,
	}
}
