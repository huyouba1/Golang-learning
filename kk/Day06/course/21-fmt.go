package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}

func time2str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	//task := &Task{
	//	id:        1,
	//	name:      "整理课程笔记",
	//	status:    0,
	//	startTime: &now,
	//	endTime:   &end,
	//	user:      "huyouba1",
	//}

	tasks := []*Task{{
		id:        1,
		name:      "整理课程笔记",
		status:    0,
		startTime: &now,
		endTime:   &end,
		user:      "huyouba1",
	},
		{
			id:        2,
			name:      "整理课程笔记2",
			status:    0,
			startTime: &now,
			endTime:   &end,
			user:      "huyouba1",
		},
	}

	file, _ := os.Create("user.txt")
	defer file.Close()
	for _, task := range tasks {
		//file.WriteString(fmt.Sprintf("%d,%s,%d,%s,%s,%s\n",
		//	task.id,
		//	task.name,
		//	task.status,
		//	time2str(task.startTime),
		//	time2str(task.endTime),
		//	task.user))

		fmt.Fprintf(file, "%d,%s,%d,%s,%s,%s\n",
			task.id,
			task.name,
			task.status,
			time2str(task.startTime),
			time2str(task.endTime),
			task.user)

		//fmt.Prin /f /ln     打印到控制台
		//fmt.Sprint /f /ln   生成字符串
		//fmt.Fprint /f /ln   输入输出流

	}

}
