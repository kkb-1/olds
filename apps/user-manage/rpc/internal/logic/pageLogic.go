package logic

import (
	"application/apps/user-manage/model"
	"application/apps/user-manage/rpc/internal/userManageXtype"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/jinzhu/copier"
	"reflect"
	"strings"

	"application/apps/user-manage/rpc/internal/svc"
	"application/apps/user-manage/rpc/userManage"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PageLogic) Page(in *userManage.UserListRequest) (*userManage.UserListResponse, error) {
	esClient := l.svcCtx.ES
	var searchBody search.Request
	var countBody count.Request
	ctx := context.Background()
	zlog := l.svcCtx.Logger.Sugar()

	if in.Query != nil {
		query := l.pushQuery(in)
		countBody.Query = &query
		searchBody.Query = &query
	}

	countOut, err := esClient.Count().Index(model.UserMangeIndex).Request(&countBody).Do(ctx)
	if err != nil {
		zlog.Errorf("count fail: %v", err)
		return nil, err
	}

	//加入排序逻辑
	l.pushSort(&searchBody)
	//加入分页逻辑
	l.pushPage(in.PageNum, in.PageSize, &searchBody)
	//分页模糊查询
	searchOut, err := esClient.Search().Index(model.UserMangeIndex).Request(&searchBody).Do(ctx)
	if err != nil {
		zlog.Errorf("search fail: %v", err)
		return nil, err
	}
	//查询出的条数
	length := searchOut.Hits.Total.Value

	var users []*userManage.User
	for _, hit := range searchOut.Hits.Hits {
		data := hit.Source_
		var user userManage.User
		var esUser model.ESUserManage
		err := json.Unmarshal(data, &esUser)
		if err != nil {
			zlog.Errorf("json fail: %v", err)
			return nil, err
		}

		err = copier.Copy(&user, &esUser)
		if err != nil {
			zlog.Errorf("copy fail: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	resp := new(userManage.UserListResponse)
	resp.List = users

	resp.TotalPage = &countOut.Count
	resp.Length = &length

	return resp, nil
}

func (l *PageLogic) pushQuery(data *userManage.UserListRequest) types.Query {
	var out types.Query
	if reflect.ValueOf(data.Query).IsZero() {
		return out
	}
	zlog := l.svcCtx.Logger.Sugar()
	var boolQuery types.BoolQuery

	var behind userManageXtype.PageQuery
	err := copier.Copy(&behind, data.Query)
	if err != nil {
		zlog.Errorf("copy错误: %v", err)
	}

	//存储query参数里的数据，零值不保存
	query := make(map[string]interface{})
	marshal, err := json.Marshal(behind)
	err = json.Unmarshal(marshal, &query)
	if err != nil {
		zlog.Errorf("json解析错误: %v", err)
	}

	var queryList []types.Query

	for k, v := range query {
		switch k {
		case "details.weight", "details.height", "details.age":
			rang := make(map[string]types.RangeQuery)
			rang[k] = v
			queryList = append(queryList, types.Query{Range: rang})

		case "details.smoke", "details.drink", "details.exercise", "parents.confirm":
			if v == 1 {
				v = true
			} else {
				v = false
			}
			term := make(map[string]types.TermQuery)
			term[k] = types.TermQuery{Value: v}
			queryList = append(queryList, types.Query{Term: term})

		case "sex":
			term := make(map[string]types.TermQuery)
			term[k] = types.TermQuery{Value: v}
			queryList = append(queryList, types.Query{Term: term})

		case "details.phone", "parents.note":
			match := make(map[string]types.MatchQuery)
			s := v.(string)
			match[k] = types.MatchQuery{Query: s}
			queryList = append(queryList, types.Query{Match: match})
		}
	}

	boolQuery.Must = queryList
	out.Bool = &boolQuery
	return out
}

func (l *PageLogic) pushSort(body *search.Request) {
	var esModel model.ESUserManage
	sort := make(map[string]string)
	tags := strings.Split(reflect.TypeOf(esModel).Field(0).Tag.Get("json"), ",")
	openId := tags[0]
	sort[openId] = "asc"
	body.Sort = append(body.Sort, sort)
}

func (l *PageLogic) pushPage(numptr, sizeptr *int64, body *search.Request) {
	var num, size int64 = 0, 10
	if numptr != nil && *numptr > 0 {
		num = *numptr
		num--
	}

	if sizeptr != nil && *sizeptr > 0 {
		size = *sizeptr
	}

	num1, size1 := int(num), int(size)

	body.From = &num1
	body.Size = &size1
}
