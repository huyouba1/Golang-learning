package main

import "fmt"

// 包级别
var packageVar string = "package Var"

func main() {
	// 函数级别
	var funcVar string = "func Var"

	{
		// 块级别
		var blockVar string = "block Var"
		fmt.Println(packageVar, funcVar, blockVar)
		{ //限定变量使用范围
			// 子块级别
			var innerVar string = "inner Var"
			fmt.Println(packageVar, funcVar, blockVar, innerVar)
		}
	}

	// 使用函数内变量
	fmt.Println(funcVar)
	// 使用包级别变量
	fmt.Println(packageVar)
}
