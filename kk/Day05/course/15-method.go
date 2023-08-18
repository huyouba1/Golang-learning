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

/*  自动生成
func (task *Task) SetName(name string) {
	task.name = name
}
*/

// user 指针接收者
func (task *Task) SetUser(user string) {
	task.user = user
}

/*  没有这个功能
func (task Task) SetUser(user string) {
	(*task).SetName(name)
}
*/

func main() {
	// 方法
	// 方法值   实例.方法名
	task := Task{}   // 值
	task2 := &Task{} // 指针
	fmt.Println(task, task2)
	// 方法表达式  结构体.方法名
	// 对于值接受者，可以通过指针也可以通过值获取方法表达式
	// GO 自动针对值接受者方法 => 自动生成指针接受者方法
	method1 := Task.SetName // (func(main.Task, string))(0x108b740)

	method1(task, "test")
	method1(*task2, "test")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	method2 := (*Task).SetName // (func(*main.Task, string))(0x108bae0)
	method2(&task, "test")
	method2(task2, "test")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	// 对于指针接受者，只能通过指针来获取方法表达式
	//method3 := Task.SetUser
	method4 := (*Task).SetUser //(func(*main.Task, string))(0x108b780)
	method4(&task, "kk")
	method4(task2, "ll")

	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)
	fmt.Printf("%#v\n", method1)
	fmt.Printf("%#v\n", method2)
	//fmt.Printf("%#v\n", method3)
	fmt.Printf("%#v\n", method4)

}
