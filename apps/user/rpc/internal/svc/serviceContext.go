package svc

import (
	"application/apps/user/rpc/internal/config"
	"application/common/xgorm"
	"application/common/xredis"
	"application/common/xzap"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Logger *zap.Logger
	DB     *gorm.DB
	Redis  *redis.Client
	//ES     *elasticsearch.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	logger := xzap.New(c.Zap)
	db := xgorm.MustOpen(c.DB, xzap.GetGormLog(logger))
	redisClient, err := xredis.New(context.Background(), c.XRedis)
	//esClient := xes.MustNew(c.ES)
	if err != nil {
		logger.Warn("redis连接失败", zap.Error(err))
	}

	return &ServiceContext{
		Config: c,
		Logger: logger,
		DB:     db,
		Redis:  redisClient,
		//ES:     esClient,
	}
}
