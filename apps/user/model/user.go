package model

import "fmt"

type User struct {
	ID        string `gorm:"primary_key" redis:"id"`
	Username  string `redis:"username"`
	Password  string `redis:"password"`
	Nickname  string `redis:"nickname"`
	Avatar    string `redis:"avatar"`
	Status    int    `redis:"status"`
	CreatedAt int    `gorm:"column:create_time" redis:"create_time"`
	UpdatedAt int    `gorm:"column:update_time" redis:"create_time"`
}

func (user *User) TableName() string {
	return "olds_user"
}

// redis前缀key
var UserPreKey = "cache:user:"

// es索引名
const UserIndex = "user_index"

func GetUsernameKey(username string) string {
	return fmt.Sprintf("%susername:%s", UserPreKey, username)
}
