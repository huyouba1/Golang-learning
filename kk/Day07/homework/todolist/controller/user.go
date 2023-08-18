package controller

import (
	"fmt"
	"os"
	"todolist/commands/config"
	"todolist/utils/ioutils"
)

// 定义退出函数
func Logout() {
	os.Exit(0)
}

// 定义登录函数，返回bool值，true则登录成功，false则登录失败退出
func Login() bool {
	for i := config.Config.LoginRetry; i > 0; i-- {
		pass := ioutils.Password("请输入密码: ")
		if pass == "hello" {
			return true
		}
		if i != 1 {
			ioutils.Error(fmt.Sprintf("密码错误，还剩余%d次机会!", i-1))
		}
	}
	ioutils.Error(fmt.Sprintf("密码错误，还剩余%d次机会!", config.Config.LoginRetry))
	return false
}
