package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// 定义Task结构体
type Task struct {
	Id        int
	Name      string
	Status    int
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

func main() {
	var tasks []*Task

	//file, err := os.Open()
	// 读取文件内容
	jsonTxt, _ := ioutil.ReadFile("task.json")
	// json反序列化
	err := json.Unmarshal(jsonTxt, &tasks)
	fmt.Println(err)

	for _, task := range tasks {
		fmt.Printf("%#v\n", task)
	}
}
