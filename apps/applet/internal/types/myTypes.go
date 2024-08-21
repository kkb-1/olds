package types

type XUserManagePage struct {
	Query    XQuery `json:"query,omitempty"`
	PageNum  int64  `json:"pageNum,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type XGetUserManageInfo struct {
	OpenId string `json:"openId,omitempty"`
	Uid    string `json:"uid,omitempty"`
}

type XQuery struct {
	Note     string      `json:"note,omitempty"`
	Confirm  int64       `json:"confirm,omitempty"`
	Phone    string      `json:"phone,omitempty"`
	Role     int64       `json:"role,omitempty"`
	Height   XFloatRange `json:"height,omitempty"`
	Weight   XFloatRange `json:"weight,omitempty"`
	Age      XIntRange   `json:"age,omitempty"`
	Sex      int64       `json:"sex,omitempty"`
	Smoke    int64       `json:"smoke,omitempty"`
	Drink    int64       `json:"drink,omitempty"`
	Exercise int64       `json:"exercise,omitempty"`
}

type XFloatRange struct {
	Gte float64 `json:"gte,omitempty"`
	Lte float64 `json:"lte,omitempty"`
}

type XIntRange struct {
	Gte int64 `json:"gte,omitempty"`
	Lte int64 `json:"lte,omitempty"`
}
