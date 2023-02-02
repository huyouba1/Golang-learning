package main

import (
	"fmt"
)

// 自定义一种类型
type Counter int
type Task map[string]string
type Callback func(...string)

func main() {
	// map[string]string{}
	// task

	// var name Type
	var cnt Counter
	fmt.Printf("%T\n", cnt)
	fmt.Printf("%#v\n", cnt)

	cnt = 1
	fmt.Printf("%#v\n", cnt)

	//var total int = 100
	//fmt.Println(total / cnt)    // 不是一种数据类型无法进行计算

	// var task map[string]string
	var task Task
	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)
	task = map[string]string{"name": "完成todolist"}
	fmt.Printf("%#v\n", task)

	// 函数类型
	var pring Callback
	pring = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, v)
		}

	}
	pring("a", "b", "c")
}
