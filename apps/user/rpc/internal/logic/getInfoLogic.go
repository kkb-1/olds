package logic

import (
	"application/apps/user/model"
	"application/apps/user/rpc/code"
	"application/common/xcache"
	"application/common/xcode"
	"context"
	"github.com/jinzhu/copier"
	"time"

	"application/apps/user/rpc/internal/svc"
	"application/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *user.GetInfoRequest) (*user.GetInfoResponse, error) {
	logger := l.svcCtx.Logger
	redisClient := l.svcCtx.Redis
	db := l.svcCtx.DB
	ctx := context.Background()
	key := model.GetUsernameKey(in.UserId)
	data := new(model.User)
	resp := new(user.GetInfoResponse)
	timout := time.Minute * 5

	data.ID = in.UserId

	err := xcache.FindByID(ctx, redisClient, db, data, timout, key, in.UserId)
	if err != nil {
		logger.Sugar().Debugf("find fail: %v", err)
		return nil, code.USERNAME_NOT_EXIST
	}

	resp.UserInfo = new(user.UserInfo)
	err = copier.Copy(resp.UserInfo, data)
	if err != nil {
		logger.Sugar().Debugf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}
	logger.Sugar().Debugf("end value of resp : %v", resp)
	return resp, nil
}
