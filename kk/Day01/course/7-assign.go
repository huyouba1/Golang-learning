package main

import "fmt"

func main() {
	var name string = "huyouba1"
	fmt.Println(name)
	name = "silence" // 更新变量的值
	fmt.Println(name)

	{
		// 不是定义,是赋值
		name = "aaaaaaaaaaa"
	}
	fmt.Println(name)
	// 输出结果为 aaaaaaaaaaa
}
