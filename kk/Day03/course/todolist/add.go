package main

import (
	"fmt"
	"strconv"
)

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "19:00", "status": statusNew, "user": "huyouba1"},
	{"id": "2", "name": "备课", "startTime": "12:00", "endTime": "1:00", "status": statusNew, "user": "huyouba1"},
	{"id": "4", "name": "复习", "startTime": "11:00", "endTime": "9:00", "status": statusNew, "user": "huyouba1"},
}

const (
	id        = "id"
	name      = "mame"
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	user      = "user"
)

const (
	statusNew     = "未执行"
	statusCompele = "完成"
)

func genId() int {
	var rt int
	for _, todo := range todos {
		todoId, _ := strconv.Atoi(todo["id"])
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}

func newTask() map[string]string {
	// id 生成(用todos中最大的ID+1)
	task := map[string]string{}
	task[id] = strconv.Itoa(genId())
	task[name] = ""
	task[startTime] = ""
	task[endTime] = ""
	task[status] = statusNew
	task[user] = ""
	return task
}

func main() {

	var text string

	task := newTask()

	fmt.Println("请输入任务信息：")
	fmt.Print("任务名：")
	fmt.Scan(&text)
	task[name] = text

	fmt.Print("开始时间：")
	fmt.Scan(&text)
	task[startTime] = text

	fmt.Print("负责人：")
	fmt.Scan(&text)
	task[user] = text

	todos = append(todos, task)
	fmt.Println(todos)
}
