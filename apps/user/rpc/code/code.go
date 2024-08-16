package code

import "application/common/xcode"

var (
	USERNAME_NOT_EXIST = xcode.New(10000, "用户不存在")
	PASSWORD_FAIL      = xcode.New(10001, "密码错误")
	USERNAME_EXIST     = xcode.New(10002, "用户已存在")
	INVITECODE_FAIL    = xcode.New(10003, "邀请码错误或过期")
)
