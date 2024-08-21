package model

// es中的聚合数据
type ESUserManage struct {
	OpenId    string       `json:"open_id,omitempty"`
	Uid       string       `json:"uid,omitempty"`
	Details   ESDetails    `json:"details,omitempty"`
	Kids      []*ESKids    `json:"kids,omitempty"`
	KidNum    int          `json:"kid_num,omitempty"`
	Parents   []*ESParents `json:"parents,omitempty"`
	ParentNum int          `json:"parent_num,omitempty"`
}

type ESDetails struct {
	Phone    string  `json:"phone,omitempty"`
	Role     int     `json:"role,omitempty"`
	Height   float64 `json:"height,omitempty"`
	Weight   float64 `json:"weight,omitempty"`
	Age      int     `json:"age,omitempty"`
	Sex      int     `json:"sex,omitempty"`
	Smoke    bool    `json:"smoke,omitempty"`
	Drink    bool    `json:"drink,omitempty"`
	Exercise bool    `json:"exercise,omitempty"`
}

type ESKids struct {
	OpenId  string `json:"open_id,omitempty"`
	Note    string `json:"note,omitempty"`
	Confirm bool   `json:"confirm,omitempty"`
}

type ESParents struct {
	Uid     string `json:"uid,omitempty"`
	Note    string `json:"note,omitempty"`
	Confirm bool   `json:"confirm,omitempty"`
}

func (doc *ESUserManage) GetDocID() string {
	return doc.OpenId
}

const (
	UserMangeIndex = "olds_user_manage"
)
