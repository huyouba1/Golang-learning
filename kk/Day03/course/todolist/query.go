package main

import (
	"fmt"
	"strings"
)

var todos = []map[string]string{
	{"id": "1", "name": "陪孩子散步", "start_time": "18:00", "end_time": "19:00", "status": statusNew, "user": "huyouba1"},
	{"id": "2", "name": "备课", "start_time": "12:00", "end_time": "1:00", "status": statusNew, "user": "huyouba1"},
	{"id": "4", "name": "复习", "start_time": "11:00", "end_time": "9:00", "status": statusNew, "user": "huyouba1"},
	{"id": "4", "name": "准备课表", "start_time": "11:00", "end_time": "9:00", "status": statusNew, "user": "huyouba1"},
	{"id": "4", "name": "课堂笔记", "start_time": "1:00", "end_time": "2:00", "status": statusNew, "user": "huyouba1"},
}

const (
	id        = "id"
	name      = "name"
	startTime = "start_time"
	endTime   = "end_time"
	status    = "status"
	user      = "user"
)

const (
	statusNew     = "未执行"
	statusCompele = "完成"
)

func printTask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID:", task[id])
	fmt.Println("任务名:", task[name])
	fmt.Println("开始时间:", task[startTime])
	fmt.Println("完成时间:", task[endTime])
	fmt.Println(strings.Repeat("-", 20))

}

func main() {

	var text string

	fmt.Print("请输入查询信息：")
	fmt.Scan(&text)

	for _, todo := range todos {
		if strings.Contains(todo[name], text) {
			printTask(todo)
		}

	}
}
