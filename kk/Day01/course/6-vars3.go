package main

import "fmt"

func main() {
	// 函数内定义变量后必须要使用
	/*
		var name string = "huyouba1"
		var mgs = "hello world"
		var desc string
	*/

	var (
		name string = "huyouba1"
		msg         = "hello world"
		desc string
	)
	x, y := "x", "y"

	fmt.Println(name, msg, desc, x, y)
}
