package main

import (
	"fmt"
	"time"
)

type User struct {
	id   int
	name string
	addr string
	tel  string
}

type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      *User
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	task = Task{
		id:   1,
		name: "完成todo",
		user: &User{
			id:   1,
			name: "acb",
		},
	}
	fmt.Printf("%#v\n", task)

	fmt.Println(task.user.id, task.user.name)
	task.user.name = "大区an"
	fmt.Println(task.user.id, task.user.name)

}
