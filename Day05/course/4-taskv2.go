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
	user      User // 组合方式，嵌入；结构体的命名组合,命名嵌入 ==》 面向对象(组合)
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	fmt.Println(task.user.name)
	task.user.name = "kk"
	fmt.Printf("%#v\n", task)

	// 赋值
	task = Task{
		id:   1,
		name: "完成todolist",
		user: User{
			id:   1,
			name: "kk",
		},
	}
	fmt.Printf("%#v\n", task)

	// 现成的user
	user := User{
		id:   1,
		name: "daquan",
	}
	task = Task{
		id:   2,
		name: "abcde",
		user: user,
	}
	fmt.Printf("%#v\n", task)

	task.user = User{id: 3, name: "chenchen"}
	fmt.Printf("%#v\n", task)

}
