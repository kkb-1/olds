package logic

import (
	"application/apps/user/model"
	"application/apps/user/rpc/internal/svc"
	"application/apps/user/rpc/user"
	"application/common/xcache"
	"application/common/xcode"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *user.UpdateRequest) (*user.UpdateResponse, error) {
	redisClient := l.svcCtx.Redis
	db := l.svcCtx.DB
	ctx := context.Background()
	data := new(model.User)
	resp := new(user.UpdateResponse)
	key := model.GetUsernameKey(in.UserId)

	err := copier.Copy(data, in)
	if err != nil {
		return nil, xcode.ServerErr
	}
	data.ID = in.UserId

	err = xcache.UpdateByID(ctx, redisClient, db, data, key)
	if err != nil {
		return nil, err
	}

	resp.UserID = in.UserId

	return resp, nil
}
