package types

import "application/apps/user-manage/model"

// kafka的detail数据
type DetailsKafka struct {
	Data []struct {
		OpenId   *string `gorm:"primary_key" json:"open_id,omitempty"`
		Uid      *string `json:"uid,omitempty"`
		Phone    *string `json:"phone,omitempty"`
		Role     *string `json:"role,omitempty"`
		Height   *string `json:"height,omitempty"`
		Weight   *string `json:"weight,omitempty"`
		Age      *string `json:"age,omitempty"`
		Sex      *string `json:"sex,omitempty"`
		Smoke    *string `json:"smoke,omitempty"`
		Drink    *string `json:"drink,omitempty"`
		Exercise *string `json:"exercise,omitempty"`
	} `json:"data"`
	Type string `json:"type"`
}

// es同步details
type DetailsES struct {
	OpenId  string           `json:"open_id,omitempty"`
	Uid     string           `json:"uid,omitempty"`
	Details *model.ESDetails `json:"details,omitempty"`
}

// kafka的binds数据
type BindsKafka struct {
	Data []struct {
		OpenId  *string `json:"open_id,omitempty"`
		Uid     *string `json:"uid,omitempty"`
		Note    *string `json:"note,omitempty"`
		Confirm *string `json:"confirm,omitempty"`
	} `json:"data"`
	Type string `json:"type"`
}

// es同步details
type BindsES model.ESParents

func IToB(value int) bool {
	if value == 1 {
		return true
	}
	return false
}
