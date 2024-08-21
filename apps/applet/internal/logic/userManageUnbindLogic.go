package logic

import (
	"application/apps/user-manage/rpc/userManage"
	"application/common/xcode"
	"context"
	"github.com/jinzhu/copier"

	"application/apps/applet/internal/svc"
	"application/apps/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserManageUnbindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserManageUnbindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserManageUnbindLogic {
	return &UserManageUnbindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserManageUnbindLogic) UserManageUnbind(req *types.UserManageUnbindRequest) (resp *types.UserManageUnbindResponse, err error) {
	rpc := l.svcCtx.UserManageRPC
	zlog := l.svcCtx.Logger.Sugar()
	in := new(userManage.UserUnbindRequest)
	err = copier.Copy(in, req)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}
	out, err := rpc.Unbind(l.ctx, in)
	if err != nil {
		zlog.Errorf("rpc fail: %v", err)
		return nil, xcode.ServerErr
	}

	resp = new(types.UserManageUnbindResponse)
	err = copier.Copy(resp, out)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}

	return
}
