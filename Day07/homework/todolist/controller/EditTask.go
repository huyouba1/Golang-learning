package controller

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"os"
	"todolist/utils/ioutils"
)

// 定义编辑任务的方法
func (c *TaskController) EditTask() {
	// 先把任务读进来
	c.SeeTask()
	id := ioutils.Input("请输入要修改的任务ID: ")
	for i := 0; i < len(TaskList); i++ {
		if TaskList[i].Id == id {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Id", "Name", "StartTime", "EndTime", "Status", "User"})
			table.Append([]string{
				TaskList[i].Id,
				TaskList[i].Name,
				time2string(TaskList[i].StartTime),
				time2string(TaskList[i].EndTime),
				TaskList[i].Status,
				TaskList[i].User,
			})
			table.Render()

			confirm := ioutils.Input("请确认任务ID y or yes: ")
			if confirm == "y" || confirm == "yes" {
				NewName := ioutils.Input("请输入新Name: ")
				TaskList[i].Name = NewName
				NewUser := ioutils.Input("请输入新User: ")
				TaskList[i].User = NewUser
				// 当文件存在的时候，就把文件清空了
				file, _ := os.Create("taskJson.json")
				defer file.Close()

				jsonEncoder := json.NewEncoder(file)
				jsonEncoder.Encode(TaskList)
				break
			} else {
				ioutils.Error("输入错误！")
				break
			}
		}
		if i == len(TaskList)-1 {
			ioutils.Error("任务ID不存在！")
		}
	}

}
