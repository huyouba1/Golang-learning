package init

import (
	"todolist/commands"
	"todolist/controller"
)

func init() {
	// 创建TaskController结构体实例
	task := controller.TaskController{}
	commands.Register("添加任务", task.AddTask)
	commands.Register("查询任务", task.SeeTask)
	commands.Register("删除任务", task.DelTask)
	commands.Register("修改任务", task.EditTask)
}
