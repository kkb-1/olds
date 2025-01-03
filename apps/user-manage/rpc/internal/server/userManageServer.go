// Code generated by goctl. DO NOT EDIT.
// Source: user-manage.proto

package server

import (
	"context"

	"application/apps/user-manage/rpc/internal/logic"
	"application/apps/user-manage/rpc/internal/svc"
	"application/apps/user-manage/rpc/userManage"
)

type UserManageServer struct {
	svcCtx *svc.ServiceContext
	userManage.UnimplementedUserManageServer
}

func NewUserManageServer(svcCtx *svc.ServiceContext) *UserManageServer {
	return &UserManageServer{
		svcCtx: svcCtx,
	}
}

func (s *UserManageServer) Page(ctx context.Context, in *userManage.UserListRequest) (*userManage.UserListResponse, error) {
	l := logic.NewPageLogic(ctx, s.svcCtx)
	return l.Page(in)
}

func (s *UserManageServer) Info(ctx context.Context, in *userManage.UserInfoRequest) (*userManage.UserInfoResponse, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}

func (s *UserManageServer) Bind(ctx context.Context, in *userManage.UserBindRequest) (*userManage.UserBindResponse, error) {
	l := logic.NewBindLogic(ctx, s.svcCtx)
	return l.Bind(in)
}

func (s *UserManageServer) Unbind(ctx context.Context, in *userManage.UserUnbindRequest) (*userManage.UserUnbindResponse, error) {
	l := logic.NewUnbindLogic(ctx, s.svcCtx)
	return l.Unbind(in)
}
