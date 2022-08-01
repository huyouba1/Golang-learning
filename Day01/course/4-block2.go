package main

import "fmt"

//var packageVar string = "package Var"

func main() {
	var packageVar string = "func Package Var"
	var funcVar string = "func Var"
	{
		var packageVar string = "block Var"
		fmt.Println("1", packageVar)
	}
	fmt.Println("2", funcVar, packageVar)
}

//子块优先级最高。相同的变量名，子块的值会覆盖父块的值
/*
1. 代码块限定变量使用范围
2. 子块可以覆盖父块中定义的变量
*/
