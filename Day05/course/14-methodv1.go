package main

// 方法值表达式

import (
	"fmt"
	"time"
)

type Task struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	user      string
}

// 如果想修改结构体里面属性的值，就用指针接收，如果不想修改结构体里面属性的值，那就用值接受者。 读的话两个都行

// name 值接收者
func (task Task) SetName(name string) {
	task.name = name
}

// user 指针接收者
func (task *Task) SetUser(user string) {
	task.user = user
}

func main() {
	// 方法
	// 方法值   实例.方法名
	task := Task{} // 值

	task2 := &Task{} // 指针

	methodValue1 := task.SetName
	methodValue2 := task2.SetName // 自动，解引用

	fmt.Printf("%#v\n", methodValue1)
	fmt.Printf("%#v\n", methodValue2)
	methodValue1("Todolist")
	methodValue2("Todolist")

	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	methodValue3 := task.SetUser
	methodValue4 := task2.SetUser // 自动，取引用
	methodValue3("kk")
	methodValue4("kk")

	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)
	// 方法表达式  结构体.方法名

}
