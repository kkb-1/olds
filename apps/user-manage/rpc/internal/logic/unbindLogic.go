package logic

import (
	"application/apps/user-manage/model"
	"context"
	"github.com/jinzhu/copier"

	"application/apps/user-manage/rpc/internal/svc"
	"application/apps/user-manage/rpc/userManage"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnbindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnbindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnbindLogic {
	return &UnbindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnbindLogic) Unbind(in *userManage.UserUnbindRequest) (*userManage.UserUnbindResponse, error) {
	var data model.Binds
	zlog := l.svcCtx.Logger.Sugar()
	db := l.svcCtx.DB

	err := copier.Copy(&data, in)
	if err != nil {
		zlog.Errorf("copy fail :%v", err)
		return nil, err
	}

	err = db.Where(data).Delete(data).Error
	if err != nil {
		zlog.Errorf("insert fail :%v", err)
		return nil, err
	}

	resp := new(userManage.UserUnbindResponse)
	err = copier.Copy(resp, in)
	if err != nil {
		zlog.Errorf("copy fail :%v", err)
		return nil, err
	}

	return resp, nil

}
