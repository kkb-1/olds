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

type DetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailsLogic {
	return &DetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailsLogic) Consume(ctx context.Context, _, value string) error {
	zlog := l.svcCtx.Logger.Sugar()
	zlog.Debugf("kafka数据:%s", value)
	//解析kafka数据
	var kmsg types.DetailsKafka
	var emsg types.DetailsES
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
	body, err := script.UpsertESDetails(emsg)
	if err != nil {
		zlog.Error(err)
	}

	//执行es
	return l.esOperate(openId, body, way)
}

func (l *DetailsLogic) esOperate(openId, body, way string) error {
	//准备es
	zlog := l.svcCtx.Logger.Sugar()
	switch way {
	case script.DELETE:
		way = "delete"
		body = ""
	default:
		way = "update"
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.ESClient,
		Index:  model.UserMangeIndex,
	})

	if err != nil {
		return err
	}

	//执行es
	err = bi.Add(l.ctx, esutil.BulkIndexerItem{
		Action:     way,
		DocumentID: openId,
		Body:       strings.NewReader(body),
		OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			zlog.Debug("details同步成功")
		},
		OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			zlog.Debug("details同步失败")
		},
	})
	if err != nil {
		return err
	}

	return bi.Close(l.ctx)
}

// kafka数据转换为es数据
func (l *DetailsLogic) kafkaToES(kmsg types.DetailsKafka, emsg *types.DetailsES) error {
	zlog := l.svcCtx.Logger.Sugar()
	var smoke, drink, exercise bool
	var height, weight float64
	var role, age, sex int
	err := copier.Copy(emsg, &(kmsg.Data[0]))
	if err != nil {
		zlog.Warn(err)
	}
	emsg.Details = new(model.ESDetails)
	err = copier.Copy(emsg.Details, &(kmsg.Data[0]))
	if err != nil {
		zlog.Warn(err)
	}

	zlog.Debugf("emsg内容：%v", *emsg)

	if kmsg.Data[0].Role != nil {
		role, _ = strconv.Atoi(*kmsg.Data[0].Role)
		emsg.Details.Role = &role
	}

	if kmsg.Data[0].Height != nil {
		height, _ = strconv.ParseFloat(*kmsg.Data[0].Height, 64)
		emsg.Details.Height = &height
	}

	if kmsg.Data[0].Weight != nil {
		weight, _ = strconv.ParseFloat(*kmsg.Data[0].Weight, 64)
		emsg.Details.Weight = &weight
	}

	if kmsg.Data[0].Age != nil {
		age, _ = strconv.Atoi(*kmsg.Data[0].Age)
		emsg.Details.Age = &age
	}

	if kmsg.Data[0].Sex != nil {
		sex, _ = strconv.Atoi(*kmsg.Data[0].Sex)
		emsg.Details.Sex = &sex

	}

	if kmsg.Data[0].Smoke != nil {
		smoke2, _ := strconv.Atoi(*kmsg.Data[0].Smoke)
		smoke = types.IToB(smoke2)
		emsg.Details.Smoke = &smoke
	}

	if kmsg.Data[0].Drink != nil {
		drink2, _ := strconv.Atoi(*kmsg.Data[0].Drink)
		drink = types.IToB(drink2)
		emsg.Details.Drink = &drink
	}

	if kmsg.Data[0].Exercise != nil {
		exercise2, _ := strconv.Atoi(*kmsg.Data[0].Exercise)
		exercise = types.IToB(exercise2)
		emsg.Details.Exercise = &exercise
	}

	zlog.Debugf("emsg内容：%v", *emsg.Details)
	return nil
}
