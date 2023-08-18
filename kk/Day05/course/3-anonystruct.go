package main

import (
	"fmt"
)

// 匿名结构体只能在函数内使用
func main() {
	// 匿名结构体  ==》 直接初始化一个变量
	//var user struct { // 类型可以省略掉
	//	id   int
	//	name string
	//	age  int
	//} = struct {
	//	id   int
	//	name string
	//	age  int
	//}{id: 1, name: "kk", age: 30}

	user := struct {
		id   int
		name string
		age  int
	}{id: 1, name: "kk", age: 30}
	fmt.Printf("%T\n", user)
	fmt.Printf("%#v\n", user)

	fmt.Println(user.name)
	user.name = "kk"
	fmt.Println(user.name)

	user = struct {
		id   int
		name string
		age  int
	}{id: 1, name: "kk", age: 30}

	fmt.Printf("%#v\n", user)

	//var name  = "1"

}
