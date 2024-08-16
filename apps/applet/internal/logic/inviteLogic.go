package logic

import (
	"application/apps/user/rpc/user"
	"application/common/jwt"
	"application/common/xcode"
	"context"
	"github.com/jinzhu/copier"

	"application/apps/applet/internal/svc"
	"application/apps/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteLogic {
	return &InviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InviteLogic) Invite() (resp *types.InviteResponse, err error) {
	userId := jwt.GetAuthValue(l.ctx)
	rpc := l.svcCtx.UserRPC
	rpcReq := new(user.InvitationRequest)
	rpcReq.UserId = userId
	invite, err := rpc.Invite(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	resp = new(types.InviteResponse)
	err = copier.Copy(resp, invite)
	if err != nil {
		return nil, xcode.ServerErr
	}

	return resp, nil
}
