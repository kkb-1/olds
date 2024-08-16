package logic

import (
	"application/common/preKey"
	"application/common/randomCode"
	"context"
	"time"

	"application/apps/user/rpc/internal/svc"
	"application/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type InviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InviteLogic {
	return &InviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InviteLogic) Invite(in *user.InvitationRequest) (*user.InvitationResponse, error) {
	code := randomCode.RandomString(4)
	key := preKey.GetInviteKey(code)
	redisClient := l.svcCtx.Redis
	timeout := time.Minute * 5

	resp := new(user.InvitationResponse)

	err := redisClient.Set(context.Background(), key, in.UserId, timeout).Err()
	if err != nil {
		return nil, err
	}

	resp.Expire = time.Now().Add(timeout).Unix()
	resp.InvitationCode = code

	return resp, nil
}
