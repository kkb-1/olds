package logic

import (
	"application/apps/user/model"
	"application/apps/user/rpc/code"
	"application/apps/user/rpc/internal/svc"
	"application/apps/user/rpc/user"
	"application/common/md5"
	"application/common/preKey"
	"application/common/snowID"
	"application/common/xcache"
	"context"
	"github.com/jinzhu/copier"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	in.Password = md5.MD5(in.Password)
	key := preKey.GetInviteKey(in.InvitationCode)
	userId, err := l.svcCtx.Redis.Get(context.Background(), key).Result()
	l.svcCtx.Redis.Del(context.Background(), key)
	if err != nil {
		return nil, code.INVITECODE_FAIL
	}

	if userId == "" {
		return nil, code.INVITECODE_FAIL
	}

	userId, err = l.RegisterOperate(in)
	if err != nil {
		return nil, err
	}
	resp := new(user.RegisterResponse)
	resp.UserId = userId
	return resp, nil
}

func (l *RegisterLogic) RegisterOperate(in *user.RegisterRequest) (string, error) {
	redisClient := l.svcCtx.Redis
	db := l.svcCtx.DB
	key := model.GetUsernameKey(in.Username)
	timeout := time.Minute * 5
	var ctx = context.Background()

	var findRedis xcache.RedisFindFunc
	var insertDb xcache.DbInsertFunc
	var expire xcache.RedisExpireFunc

	findRedis = func() (bool, error) {
		result, err := redisClient.Exists(ctx, key).Result()
		if result == 0 {
			return false, err
		}
		return true, code.USERNAME_EXIST
	}
	data := new(model.User)

	insertDb = func() error {
		err := copier.Copy(data, in)
		data.ID, err = snowID.GetID(snowID.USER_NODE)
		data.Status = model.UserStatusNormal
		if err != nil {
			return err
		}

		err = db.Create(data).Error
		if err != nil {
			return code.USERNAME_EXIST
		}
		return nil
	}

	expire = func() {
		redisClient.Expire(ctx, key, timeout)
	}

	err := xcache.InsertByCache(ctx, findRedis, insertDb, expire)
	if err != nil {
		return "", err
	}

	return data.ID, nil
}
