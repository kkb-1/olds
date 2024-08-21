package logic

import (
	"application/apps/user-manage/model"
	"application/apps/user-manage/rpc/internal/svc"
	"application/apps/user-manage/rpc/userManage"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *userManage.UserInfoRequest) (*userManage.UserInfoResponse, error) {
	logx.Info(777777)
	if in.OpenId != nil {
		return l.findByOpenId(*in.OpenId)
	}
	if in.Uid != nil {
		return l.findByUid(*in.Uid)
	}

	return &userManage.UserInfoResponse{}, nil
}

func (l *InfoLogic) findByOpenId(openId string) (*userManage.UserInfoResponse, error) {
	zlog := l.svcCtx.Logger.Sugar()
	var esUser model.ESUserManage
	var user userManage.User
	esClient := l.svcCtx.ES
	ctx := context.Background()
	out, err := esClient.Get(model.UserMangeIndex, openId).Do(ctx)
	if err != nil {
		zlog.Errorf("get fail: %v", err)
		return nil, err
	}

	data := out.Source_

	err = json.Unmarshal(data, &esUser)
	if err != nil {
		zlog.Errorf("json fail: %v", err)
		return nil, err
	}

	err = copier.Copy(&user, &esUser)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, err
	}

	resp := new(userManage.UserInfoResponse)
	resp.Info = &user

	return resp, nil
}

func (l *InfoLogic) findByUid(uid string) (*userManage.UserInfoResponse, error) {
	zlog := l.svcCtx.Logger.Sugar()
	var esUser model.ESUserManage
	var user userManage.User
	esClient := l.svcCtx.ES
	ctx := context.Background()
	var query types.Query

	query.Term["uid"] = types.TermQuery{Value: uid}

	out, err := esClient.Search().Request(&search.Request{Query: &query}).Do(ctx)
	if err != nil {
		zlog.Errorf("search fail: %v", err)
		return nil, err
	}

	data := out.Hits.Hits[0].Source_

	err = json.Unmarshal(data, &esUser)
	if err != nil {
		zlog.Errorf("json fail: %v", err)
		return nil, err
	}

	err = copier.Copy(&user, &esUser)
	if err != nil {
		zlog.Errorf("copy fail: %v", err)
		return nil, err
	}

	resp := new(userManage.UserInfoResponse)
	resp.Info = &user

	return resp, nil
}
