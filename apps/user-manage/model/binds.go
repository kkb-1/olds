package model

type Binds struct {
	OpenId  *string `gorm:"primary_key" json:"open_id,omitempty"`
	Uid     *string `gorm:"primary_key" json:"uid,omitempty"`
	Note    *string `json:"note,omitempty"`
	Confirm *int    `json:"confirm,omitempty"`
}

func (b *Binds) TableName() string {
	return "binds"
}
