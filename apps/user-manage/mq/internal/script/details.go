package script

import (
	"application/apps/user-manage/mq/internal/types"
	"encoding/json"
	"fmt"
)

const (
	//同步details数据的脚本
	UpsertESDetailsScript = `{"doc":%s, "doc_as_upsert":true}`
)

func UpsertESDetails(es types.DetailsES) (string, error) {
	marshal, err := json.Marshal(es)
	if err != nil {
		return "", err
	}
	str := string(marshal)
	str = fmt.Sprintf(UpsertESDetailsScript, str)
	return str, nil
}
