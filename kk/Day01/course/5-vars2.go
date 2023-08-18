package main

import "fmt"

func main() {
	var name string = "huyouba1" //定义了类型并且初始化了值

	var zeroString string // 定义了变量类型，但没初始化值

	var typeString = "huyouba1" //定义变量省略类型，不能省略初始化值
	// 通过对应的值类型推导变量类型

	//短声明(必须在函数内使用子块使用，不能在包级别使用)
	shortString := "huyouba1"
	// 通过对应的值类型推导变量类型

	// 初始化使用类型对应的零值（空字符串）
	fmt.Println(name, zeroString)
	fmt.Println(typeString)
	fmt.Println(shortString)
}
