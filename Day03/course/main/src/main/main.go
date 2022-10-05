package main

import "fmt"

var mainVar = "main Var"

func mainFunc() {
	fmt.Println("main Func")

}
func main() {
	mainFunc()
	fmt.Println(mainVar)
	utilsFunc()
	fmt.Println(utilsVar)
}

// GO PATH 项目

// GO MODULE 项目
// GO包
// 1. 同一个文件夹下的所有go文件的包名必须一致
// 2. 关闭了 GOMODULE
// GOPATH 在项目目录直接运行 go build 无文件名
// 将当前文件夹下的所有 go 文件进行编译
// 3. main 包编译为可执行程序
// 4. main 包里面只能有一个 main 函数

// GOPATH 环境变量信息，定义多个目录
// src ==> 源文件
// pkg ==> 程序创建的包文件
// bin ==> 程序编译的可执行文件

// 编译程序 go build 项目文件路径(src)
// 导入包使用包的路径名
// 调用方法/数据，使用包名/变量名/函数名
