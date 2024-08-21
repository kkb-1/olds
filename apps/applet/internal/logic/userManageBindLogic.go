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

type UserManageBindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserManageBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserManageBindLogic {
	return &UserManageBindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserManageBindLogic) UserManageBind(req *types.UserManageBindRequest) (resp *types.UserManageBindResponse, err error) {
	rpc := l.svcCtx.UserManageRPC
	zlog := l.svcCtx.Logger.Sugar()
	in := new(userManage.UserBindRequest)
	err = copier.Copy(in, req)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}
	out, err := rpc.Bind(l.ctx, in)
	if err != nil {
		zlog.Errorf("rpc fail: %v", err)
		return nil, xcode.ServerErr
	}

	resp = new(types.UserManageBindResponse)
	err = copier.Copy(resp, out)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}

	return
}
