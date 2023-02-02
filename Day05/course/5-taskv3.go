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
	User      // 组合方式，嵌入(匿名嵌入) =》 面向对象(继承) 也是有一个属性名字(默认简写)
	Creator
}

type Creator struct {
	id   int
	name string
	addr string
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)
	fmt.Println(task.User.name)
	task.User.name = "abcdasdas"
	fmt.Printf("%#v\n", task)

	// 赋值
	task = Task{
		id:   1,
		name: "完成todolist",
		User: User{
			id:   1,
			name: "daqruan",
			addr: "xiadian",
		},
	}
	fmt.Printf("%#v\n", task)

	fmt.Println(task.name)
	fmt.Println(task.Creator.addr)

	task.Creator.addr = "xinjiapo"
	fmt.Println(task.Creator.addr)
	fmt.Printf("%#v\n", task)

	task.User.addr = "aisaiebiya"
	fmt.Println(task.User.addr)
	fmt.Printf("%#v\n", task)

}
