package main

import "todolist/commands"

// 导入初始化包，但是在main函数中不适用init包中的方法，只是在此处调用init包中的init方法。使用 _ 导入
import _ "todolist/init"

func main() {
	// 程序入口
	commands.Run()
}
