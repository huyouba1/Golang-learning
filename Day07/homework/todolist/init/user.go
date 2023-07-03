package init

import (
	"todolist/commands"
	"todolist/controller"
)

// 对user进行初始化
// 初始化函数会在导入包(init包)的时候执行

func init() {
	// 传入Name和CallBack 然后注册到了mgr中
	commands.Register("退出", controller.Logout)

	// 传入Login函数，注册到mgr中
	commands.RegisterLoginCallBack(controller.Login)
}
