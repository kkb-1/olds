package logic

import (
	"application/apps/user/rpc/user"
	"application/common/jwt"
	"application/common/xcode"
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"

	"application/apps/applet/internal/svc"
	"application/apps/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResponse, err error) {
	userId := jwt.GetAuthValue(l.ctx)
	rpc := l.svcCtx.UserRPC
	logger := l.svcCtx.Logger

	rpcReq := new(user.GetInfoRequest)
	rpcReq.UserId = userId
	info, err := rpc.GetInfo(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	resp = new(types.GetUserInfoResponse)
	err = copier.Copy(&resp.UserInfo, info.UserInfo)
	if err != nil {
		logger.Error("copy fail", zap.Error(err))
		return nil, xcode.ServerErr
	}

	return
}
