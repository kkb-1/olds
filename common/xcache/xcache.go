package xcache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

//该包宗旨是设计缓存策略，而具体实施的sql语句等等并不是该包重点，所以不做干预，提高自由度

// 存在返回true，不存在返回false
type RedisFindFunc func() (exist bool, err error)

type RedisExpireFunc func()

type RedisSetFunc func() error

type RedisDeleteFunc func() error

type DbFindFunc func() error

type DbUpdateFunc func() error

type DbInsertFunc func() error

type EmptyCacheSetFunc func()

func FindByCache(ctx context.Context, findRedis RedisFindFunc, findDB DbFindFunc, setRedis RedisSetFunc, expire RedisExpireFunc, emptyCache EmptyCacheSetFunc) error {
	expire()
	exist, err := findRedis()
	if !exist {
		err = findDB()
		if err != nil {

			//空缓存策略
			emptyCache()

			return err
		}

		err = setRedis()
		expire()
	}

	return nil
}

func UpdateByCache(ctx context.Context, deleteRedis RedisDeleteFunc, updateDB DbUpdateFunc) error {
	err := deleteRedis()

	err = updateDB()
	if err != nil {
		return err
	}

	err = deleteRedis()
	return nil
}

// 这种策略适用于需要先确定是否存在的情况
func InsertByCache(ctx context.Context, findRedis RedisFindFunc, insertDB DbInsertFunc, expire RedisExpireFunc) error {
	expire()
	exist, err := findRedis()
	if !exist {
		err = insertDB()
		if err != nil {
			return err
		}
	}
	return err
}

// 该方法适合通过id查询全部信息的情况
func FindByID(ctx context.Context, redisClient *redis.Client, db *gorm.DB, model interface{}, timeout time.Duration, key, id string) error {
	var findRedis RedisFindFunc
	var findDb DbFindFunc
	var setRedis RedisSetFunc
	var expire RedisExpireFunc
	var empty EmptyCacheSetFunc

	findRedis = func() (exist bool, err error) {
		result, err := redisClient.Exists(ctx, key).Result()
		if err != nil {
			return false, err
		}
		if result <= 0 {
			return false, nil
		}

		err = redisClient.HGetAll(ctx, key).Scan(model)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	findDb = func() error {
		err := db.Model(model).First(model).Error
		return err
	}

	setRedis = func() error {
		err := redisClient.HSet(ctx, key, model).Err()
		return err
	}

	empty = func() {
		redisClient.Set(ctx, key, "*", timeout)
	}

	expire = func() {
		redisClient.Expire(ctx, key, timeout)
	}

	return FindByCache(ctx, findRedis, findDb, setRedis, expire, empty)
}

func UpdateByID(ctx context.Context, redisClient *redis.Client, db *gorm.DB, model interface{}, key string) error {
	var deleteRedis RedisDeleteFunc
	var updateDb DbUpdateFunc

	deleteRedis = func() error {
		return redisClient.Del(ctx, key).Err()
	}

	updateDb = func() error {
		return db.Model(model).Updates(model).Error
	}

	return UpdateByCache(ctx, deleteRedis, updateDb)
}
