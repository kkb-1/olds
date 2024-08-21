package logic

import (
	"application/apps/user-manage/rpc/userManage"
	"application/common/xcode"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"application/apps/applet/internal/svc"
	"application/apps/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserManagePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserManagePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserManagePageLogic {
	return &UserManagePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserManagePageLogic) UserManagePage(req *types.UserManagePageRequest) (resp *types.UserManagePageResponse, err error) {
	rpc := l.svcCtx.UserManageRPC
	zlog := l.svcCtx.Logger.Sugar()
	in := new(userManage.UserListRequest)
	behind := new(types.XUserManagePage)
	err = copier.Copy(behind, req)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}
	marshal, err := json.Marshal(behind)
	err = json.Unmarshal(marshal, in)
	if err != nil {
		zlog.Errorf("json fail: %v", err)
		return nil, xcode.ServerErr
	}

	out, err := rpc.Page(l.ctx, in)
	if err != nil {
		zlog.Errorf("rpc fail: %v", err)
		return nil, xcode.ServerErr
	}
	zlog.Debugf("out: %v", out)
	resp = new(types.UserManagePageResponse)
	err = copier.Copy(resp, out)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}

	return
}
