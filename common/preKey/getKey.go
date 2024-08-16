package preKey

import "fmt"

//用于注册私人的完整key

// 邀请码key
func GetInviteKey(code string) string {
	return fmt.Sprintf("%sinvite:%s", CODE, code)
}
