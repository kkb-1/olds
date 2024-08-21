package logic

import (
	"application/apps/user-manage/model"
	"context"
	"github.com/jinzhu/copier"

	"application/apps/user-manage/rpc/internal/svc"
	"application/apps/user-manage/rpc/userManage"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindLogic {
	return &BindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindLogic) Bind(in *userManage.UserBindRequest) (*userManage.UserBindResponse, error) {
	var data model.Binds
	var confirm int = 1
	zlog := l.svcCtx.Logger.Sugar()
	db := l.svcCtx.DB

	err := copier.Copy(&data, in)
	if err != nil {
		zlog.Errorf("copy fail :%v", err)
		return nil, err
	}
	data.Confirm = &confirm

	err = db.Create(data).Error
	if err != nil {
		zlog.Errorf("insert fail :%v", err)
		return nil, err
	}

	resp := new(userManage.UserBindResponse)
	err = copier.Copy(resp, in)
	if err != nil {
		zlog.Errorf("copy fail :%v", err)
		return nil, err
	}

	return resp, nil
}
