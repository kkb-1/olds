package logic

import (
	"application/apps/user/model"
	"application/apps/user/rpc/code"
	"application/apps/user/rpc/internal/svc"
	"application/apps/user/rpc/user"
	"application/common/md5"
	"application/common/xcache"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	data, err := l.FindByUsername(in.Username)
	if err != nil {
		return nil, err
	}

	if md5.MD5(in.Password) != data.Password {
		return nil, code.PASSWORD_FAIL
	}

	u := new(user.LoginResponse)
	u.UserId = data.ID

	return u, nil
}

func (l *LoginLogic) FindByUsername(username string) (*model.User, error) {

	data := new(model.User)

	var ctx = context.Background()

	var findRedis xcache.RedisFindFunc
	var findDb xcache.DbFindFunc
	var setRedis xcache.RedisSetFunc
	var expire xcache.RedisExpireFunc
	var empty xcache.EmptyCacheSetFunc

	redisClient := l.svcCtx.Redis
	db := l.svcCtx.DB

	key := model.GetUsernameKey(username)
	timeout := time.Minute * 5

	findRedis = func() (bool, error) {
		result, err := redisClient.Exists(ctx, key).Result()
		if result == 0 {
			return false, nil
		}
		err = redisClient.HGetAll(ctx, key).Scan(data)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	findDb = func() error {
		err := db.Model(&model.User{}).
			Where("username = ?", username).
			Not("status = ?", model.UserStatusBan).
			First(data).Error

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return code.USERNAME_NOT_EXIST
		}

		return err
	}

	setRedis = func() error {
		err := redisClient.HSet(ctx, key, data).Err()
		if err != nil {
			l.svcCtx.Logger.Error("setRedis,err:", zap.Error(err))
		}
		return err
	}

	expire = func() {
		redisClient.Expire(ctx, key, timeout)
	}

	empty = func() {
		redisClient.Set(ctx, key, "*", time.Minute)
	}

	return data, xcache.FindByCache(ctx, findRedis, findDb, setRedis, expire, empty)

}
