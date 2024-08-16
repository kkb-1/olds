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

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	rpc := l.svcCtx.UserRPC
	logger := l.svcCtx.Logger

	rpcReq := new(user.RegisterRequest)
	err = copier.Copy(rpcReq, req)
	if err != nil {
		logger.Error("copy fail", zap.Error(err))
		return nil, xcode.ServerErr
	}

	register, err := rpc.Register(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	access := l.svcCtx.Config.JWT
	token, err := jwt.GetAuthToken(access.AccessSecret, access.AccessExpire, register.UserId)
	if err != nil {
		logger.Error("token get fail", zap.Error(err))
		return nil, xcode.ServerErr
	}

	resp = new(types.RegisterResponse)
	err = copier.Copy(&resp.Token, token)
	if err != nil {
		logger.Error("copy fail", zap.Error(err))
		return nil, xcode.ServerErr
	}

	return resp, nil
}
