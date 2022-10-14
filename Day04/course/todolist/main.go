package main

import (
	"fmt"
	"todolist/controller/task"
	mtask "todolist/models/task"
	// 1. 导入mod init name + 目录结构
	// 2. 包名与所在目录名称保持一致
	// 3. 调用时用 包名.变量名
	"github.com/huyouba1/strutil"
)

func main() {
	fmt.Println(mtask.Name)
	fmt.Println(task.Name)
	task.Call()
	fmt.Println(task.Version)
	task.PrintVersion()
	task.GetVersion()
	fmt.Println(strutil.RandString())
}
