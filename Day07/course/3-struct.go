package main

import "fmt"

type Task struct {
}

// 值接受者
func (t Task) Value() {
	fmt.Println("value")
}

// 指针接受者
func (t *Task) PValue() {
	fmt.Println("PValue")
}

func main() {
	var task Task
	task.Value()
	task.PValue() // task没有PValue方法， 语法糖
	(&task).PValue()

	var PTask *Task = new(Task)
	(*PTask).PValue() //
	PTask.Value()     // ptask 有 value 方法，语法糖
	PTask.PValue()
}
