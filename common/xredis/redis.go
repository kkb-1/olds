package xredis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// 获取redis客户端
func New(ctx context.Context, config Config) (*redis.Client, error) {

	opt := &redis.Options{
		Addr:     getAdder(config),
		Password: config.Password,
		DB:       config.DB,
	}

	client := redis.NewClient(opt)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func MustNew(ctx context.Context, config Config) *redis.Client {
	cilent, err := New(ctx, config)
	if err != nil {
		panic(err)
	}

	return cilent
}

func getAdder(config Config) string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
