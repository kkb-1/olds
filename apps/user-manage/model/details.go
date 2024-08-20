package model

type Details struct {
	OpenId   *string  `gorm:"primary_key" json:"open_id,omitempty"`
	Uid      *string  `json:"uid,omitempty"`
	Phone    *string  `json:"phone,omitempty"`
	Role     *int     `json:"role,omitempty"`
	Height   *float64 `json:"height,omitempty"`
	Weight   *float64 `json:"weight,omitempty"`
	Age      *int     `json:"age,omitempty"`
	Sex      *int     `json:"sex,omitempty"`
	Smoke    *int     `json:"smoke,omitempty"`
	Drink    *int     `json:"drink,omitempty"`
	Exercise *int     `json:"exercise,omitempty"`
}
