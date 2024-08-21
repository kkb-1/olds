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

type GetUserManageInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserManageInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserManageInfoLogic {
	return &GetUserManageInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserManageInfoLogic) GetUserManageInfo(req *types.GetUserManageInfoRequest) (resp *types.GetUserManageInfoResponse, err error) {
	rpc := l.svcCtx.UserManageRPC
	zlog := l.svcCtx.Logger.Sugar()
	in := new(userManage.UserInfoRequest)
	behind := new(types.XGetUserManageInfo)

	err = copier.Copy(behind, req)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}

	marshal, err := json.Marshal(behind)
	err = json.Unmarshal(marshal, in)
	zlog.Debugf("in: %v", in)
	if err != nil {
		zlog.Errorf("json fail: %v", err)
		return nil, err
	}

	out, err := rpc.Info(l.ctx, in)
	if err != nil {
		zlog.Errorf("rpc fail: %v", err)
		return nil, xcode.ServerErr
	}

	resp = new(types.GetUserManageInfoResponse)
	err = copier.Copy(resp, out)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, xcode.ServerErr
	}

	return
}
