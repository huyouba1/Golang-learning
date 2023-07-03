package controller

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"os"
	"todolist/utils/ioutils"
)

// 定义删除任务的方法
func (c *TaskController) DelTask() {
	// 先读取任务到TaskList中
	c.SeeTask()
	id := ioutils.Input("请输入要删除的任务ID: ")
	// 打印出用户选择的任务，让用户确认
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
				// 删除指定ID的task
				TaskList = append(TaskList[:i], TaskList[i+1:]...)
				// 当文件存在时，则清空文件内容
				file, _ := os.Create("taskJson.json")
				defer file.Close()
				jsonEncoder := json.NewEncoder(file)
				jsonEncoder.Encode(TaskList)
				break
			} else {
				ioutils.Error("输入错误")
				break
			}
		}
		if i == len(TaskList)-1 {
			ioutils.Error("任务ID不存在")
		}
	}
}
