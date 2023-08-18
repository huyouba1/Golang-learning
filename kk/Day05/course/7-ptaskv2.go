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
	*User
	user2 User
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)
	task = Task{
		id:   1,
		name: "完成todo",
		User: &User{
			id:   1,
			name: "asda",
			addr: "Beijing",
		},
		user2: User{
			id:   1,
			name: "asda",
			addr: "Beijing",
		},
	}
	task2 := task
	task2.name = "完成整理"
	fmt.Println(task)
	fmt.Println(task2)

	task.User.addr = " 北京" // 重新赋值之后，由于是指针类型，所以两个地方都会受影响
	fmt.Println(task.User)
	fmt.Println(task2.User)

	task.user2.addr = " 北京"
	fmt.Println(task.user2)
	fmt.Println(task2.user2)
}
