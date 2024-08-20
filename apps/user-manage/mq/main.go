package main

import (
	"application/apps/user-manage/mq/internal/config"
	"application/apps/user-manage/mq/internal/logic"
	"application/apps/user-manage/mq/internal/svc"
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

// 目的是借助kafka自主消费消息，为es同步数据
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	err := c.ServiceConf.SetUp()
	if err != nil {
		panic(err)
	}

	logx.DisableStat()
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(mq)
	}
	fmt.Println("mq服务启动成功")
	serviceGroup.Start()
}
