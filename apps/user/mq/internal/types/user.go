package types

type UserCanalMsg struct {
	Data []struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
		Status   string `json:"status"`
		CreateAt string `json:"create_time"`
		UpdateAt string `json:"update_time"`
	} `json:"data"`
}

type UserESMsg struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Status   int    `json:"status"`
	CreateAt int    `json:"create_time"`
	UpdateAt int    `json:"update_time"`
}
