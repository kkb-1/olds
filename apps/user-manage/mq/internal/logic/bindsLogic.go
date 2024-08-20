package logic

import (
	"application/apps/user-manage/model"
	"application/apps/user-manage/mq/internal/script"
	"application/apps/user-manage/mq/internal/svc"
	"application/apps/user-manage/mq/internal/types"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

type BindsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindsLogic {
	return &BindsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindsLogic) Consume(ctx context.Context, _, value string) error {
	zlog := l.svcCtx.Logger.Sugar()
	zlog.Debugf("kafka数据:%s", value)
	//解析kafka数据
	var kmsg types.BindsKafka
	var emsg types.BindsES
	err := json.Unmarshal([]byte(value), &kmsg)
	if err != nil {
		zlog.Error(err)
	}
	openId := *kmsg.Data[0].OpenId
	way := kmsg.Type

	//将kafka数据转换成es数据
	err = l.kafkaToES(kmsg, &emsg)
	if err != nil {
		zlog.Error(err)
	}

	//获取脚本
	body, err := script.SyncBinds(emsg, way)
	if err != nil {
		zlog.Error(err)
	}

	//执行es
	return l.esOperate(openId, body)
}

func (l *BindsLogic) esOperate(openId, body string) error {
	//准备es
	zlog := l.svcCtx.Logger.Sugar()

	zlog.Debugf("open_id:%s, body: %s", openId, body)

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.ESClient,
		Index:  model.UserMangeIndex,
	})
	if err != nil {
		return err
	}

	//执行es
	err = bi.Add(l.ctx, esutil.BulkIndexerItem{
		Action:     "update",
		DocumentID: openId,
		Body:       strings.NewReader(body),
		OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			zlog.Debug("binds同步成功")
		},
		OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			zlog.Debug("binds同步失败")
		},
	})
	if err != nil {
		return err
	}

	return bi.Close(l.ctx)
}

// kafka数据转换为es数据
func (l *BindsLogic) kafkaToES(kmsg types.BindsKafka, emsg *types.BindsES) error {
	zlog := l.svcCtx.Logger.Sugar()
	var confirm bool
	err := copier.Copy(emsg, &(kmsg.Data[0]))
	if err != nil {
		zlog.Warn(err)
	}
	zlog.Debugf("emsg内容：%v", *emsg)
	if kmsg.Data[0].Confirm != nil {
		con, err := strconv.Atoi(*kmsg.Data[0].Confirm)
		if err != nil {
			zlog.Error(err)
		}
		confirm = types.IToB(con)
		emsg.Confirm = &confirm
	}

	zlog.Debugf("emsg内容：%v", *emsg)
	return nil
}
