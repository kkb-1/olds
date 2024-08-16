package randomCode

import (
	"math/rand"
	"time"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func RandomString(length int) string {
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	code := make([]byte, length)

	for i, _ := range code {
		code[i] = charSet[seed.Intn(len(charSet))]
	}

	return string(code)
}
