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

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {
	userId := jwt.GetAuthValue(l.ctx)

	rpc := l.svcCtx.UserRPC
	logger := l.svcCtx.Logger

	rpcReq := new(user.UpdateRequest)
	err = copier.Copy(rpcReq, req)
	if err != nil {
		logger.Error("copy fail", zap.Error(err))
		return nil, xcode.ServerErr
	}
	rpcReq.UserId = userId
	logger.Sugar().Debugf("rpcReq value: %v", rpcReq)

	info, err := rpc.UpdateInfo(l.ctx, rpcReq)
	logger.Sugar().Debugf("rpc userinfo value: %v", info)
	if err != nil {
		return nil, err
	}

	resp = new(types.UpdateUserInfoResponse)
	resp.UserId = info.UserID

	return resp, nil
}
