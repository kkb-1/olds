package script

import (
	"application/apps/user-manage/mq/internal/types"
	"encoding/json"
	"fmt"
)

const (
	Parents   = `parents`
	ParentNum = `parent_num`
	Uid       = `uid`

	//插入binds数据脚本
	InsertBindsScript = `{"script" : {"source": "if(ctx._source.%s==null){ctx._source.%s=[params.obj];}else{ctx._source.%s.add(params.obj);}ctx._source.%s = ctx._source.%s.length;","lang": "painless","params" : {"obj" : %s}},"upsert": {"parents":[],"parent_num": 0}}`

	//删除binds数据脚本
	DeleteBindsScript = `{"script" : {"source": "for (def i = 0; i < ctx._source.%s.length; i++) {if (ctx._source.%s[i].%s == params.%s){ctx._source.%s.remove(i);ctx._source.%s = ctx._source.%s.length;break}}","params" : %s}}`

	//更新binds数据脚本
	UpdateBindsScript = `{"script" : {"source": "for (def i = 0; i < ctx._source.%s.length; i++) {if (ctx._source.%s[i].%s == params.%s){for (def j = 0; j < params.fields.length; j++)ctx._source.%s[i][params.fields[j]] = params.values[j];break}}","lang": "painless","params" : %s}}`
)

func SyncBinds(es types.BindsES, way string) (string, error) {
	value, err := json.Marshal(es)
	str := string(value)
	if err != nil {
		return "", err
	}

	switch way {
	case INSERT:
		str = fmt.Sprintf(InsertBindsScript, Parents, Parents, Parents, ParentNum, Parents, str)
	case DELETE:
		str = fmt.Sprintf(DeleteBindsScript, Parents, Parents, Uid, Uid, Parents, ParentNum, Parents, str)
	case UPDATE:
		str = fmt.Sprintf(UpdateBindsScript, Parents, Parents, Uid, Uid, Parents, updateBindsParams(es))
	}

	return str, nil
}

// 真的逆天，但是我不想用反射，因为字段不是很多，而且基本不会变动，无论是性能还是不易出错来看，这样子都是最方便的
func updateBindsParams(es types.BindsES) string {
	var fields string = `[`
	var values string = `[`

	if es.Note != nil {
		fields = fields + `"note"`
		values = fmt.Sprintf(`%s"%v"`, values, *es.Note)
	}

	if fields != `[` {
		fields += `,`
		values += `,`
	}
	if es.Confirm != nil {
		fields = fields + `"confirm"`
		values = fmt.Sprintf(`%s%v`, values, *es.Confirm)
	}

	return fmt.Sprintf(`{"uid":"%s", "fields":%s],"values":%s]}`, *es.Uid, fields, values)
}
