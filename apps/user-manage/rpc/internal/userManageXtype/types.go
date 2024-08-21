package userManageXtype

import "application/apps/user-manage/rpc/userManage"

// 顺应es语法
type PageQuery struct {
	Note     string                 `json:"parents.note,omitempty"`
	Confirm  int64                  `json:"parents.confirm,omitempty"`
	Phone    string                 `json:"details.phone,omitempty"`
	Role     int64                  `json:"details.role,omitempty"`
	Height   *userManage.FloatRange `json:"details.height,omitempty"`
	Weight   *userManage.FloatRange `json:"details.weight,omitempty"`
	Age      *userManage.IntRange   `json:"details.age,omitempty"`
	Sex      int64                  `json:"details.sex,omitempty"`
	Smoke    int64                  `json:"details.smoke,omitempty"`
	Drink    int64                  `json:"details.drink,omitempty"`
	Exercise int64                  `json:"details.exercise,omitempty"`
}
