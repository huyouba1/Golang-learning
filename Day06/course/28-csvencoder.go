package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Task struct {
	Id        int
	Name      string
	Status    int
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

// 定义常量 时间格式字符串
const TimeLayout = "2006-01-02 15:04:05"

func time2str(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(TimeLayout)
}

func main() {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	tasks := []*Task{{
		Id:        1,
		Name:      "整理课程笔记",
		Status:    0,
		StartTime: &now,
		EndTime:   &end,
		User:      "huyouba1",
	},
		{
			Id:        2,
			Name:      "整理课程笔记2",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "huyouba1",
		},
	}

	// 创建文件
	file, _ := os.Create("user.csv")
	defer file.Close()
	// 创建带缓冲IO的写对象
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// 创建csv写对象
	csvWriter := csv.NewWriter(writer)
	// 写入csv头
	csvWriter.Write([]string{"ID", "任务名称", "状态", "开始时间", "结束时间", "执行者"})
	for _, task := range tasks {
		csvWriter.Write([]string{
			strconv.Itoa(task.Id),
			task.Name,
			strconv.Itoa(task.Status),
			time2str(task.StartTime),
			time2str(task.EndTime),
			task.User,
		})
	}
}
