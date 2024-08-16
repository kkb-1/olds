package logic

import (
	"application/apps/user/model"
	"application/apps/user/mq/internal/svc"
	"application/apps/user/mq/internal/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
)

type UserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLogic) Consume(ctx context.Context, _, value string) error {
	logx.Debugf("userKafka消费数据:%s", value)
	var msg *types.UserCanalMsg
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		logx.Errorf("解析json失败，err：%v", err)
		return err
	}
	return nil
}

// 将canal格式数据转化为es格式
func (l *UserLogic) userOperate(msg *types.UserCanalMsg) error {
	if len(msg.Data) == 0 {
		return nil
	}

	var esData []*types.UserESMsg
	for _, data := range msg.Data {
		status, _ := strconv.Atoi(data.Status)
		createTime, _ := strconv.Atoi(data.CreateAt)
		updateTime, _ := strconv.Atoi(data.UpdateAt)

		esMsg := new(types.UserESMsg)
		err := copier.Copy(esMsg, msg)
		if err != nil {
			logx.Errorf("copy失败,err:%v", err)
		}

		esMsg.Status = status
		esMsg.CreateAt = createTime
		esMsg.UpdateAt = updateTime

		esData = append(esData, esMsg)
	}

	err := l.BatchUpsertToES(l.ctx, esData)
	if err != nil {
		l.Logger.Errorf("BatchUpSertToEs data: %v error: %v", esData, err)
	}

	return err
}

// 同步数据到es
func (l *UserLogic) BatchUpsertToES(ctx context.Context, data []*types.UserESMsg) error {
	if len(data) == 0 {
		return nil
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client: l.svcCtx.ESClient,
		Index:  model.UserIndex,
	})
	if err != nil {
		return err
	}

	for _, d := range data {
		v, err := json.Marshal(d)
		if err != nil {
			return err
		}

		payload := fmt.Sprintf(`{"doc":%s,"doc_as_upsert":true}`, string(v))
		err = bi.Add(ctx, esutil.BulkIndexerItem{
			Action:     "update",
			DocumentID: d.ID,
			Body:       strings.NewReader(payload),
			OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
			},
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem, err error) {
			},
		})
		if err != nil {
			return err
		}
	}

	return bi.Close(ctx)
}
