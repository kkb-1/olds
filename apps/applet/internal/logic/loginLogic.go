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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	logger := l.svcCtx.Logger
	rpc := l.svcCtx.UserRPC

	rpcReq := new(user.LoginRequest)
	err = copier.Copy(rpcReq, req)
	if err != nil {
		logger.Error("拷贝失败", zap.Error(err))
		return nil, xcode.ServerErr
	}

	login, err := rpc.Login(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	access := l.svcCtx.Config.JWT

	token, err := jwt.GetAuthToken(access.AccessSecret, access.AccessExpire, login.UserId)
	if err != nil {
		logger.Error("token生成失败", zap.Error(err))
		return nil, xcode.ServerErr
	}

	resp = new(types.LoginResponse)
	err = copier.Copy(&resp.Token, token)
	if err != nil {
		logger.Error("拷贝失败", zap.Error(err))
		return nil, xcode.ServerErr
	}

	return resp, nil
}
