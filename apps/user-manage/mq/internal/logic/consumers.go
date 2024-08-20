package logic

import (
	"application/apps/user-manage/mq/internal/svc"
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// 为kafka注册并分配任务
func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.DetailsKqConsumer, NewDetailsLogic(ctx, svcCtx)),
		kq.MustNewQueue(svcCtx.Config.BindsKqConsumer, NewBindsLogic(ctx, svcCtx)),
	}
}
