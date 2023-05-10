package main

import (
	"fmt"
	"os"
)

func main() {
	os.Stdout.WriteString("我是kk") // 运行程序命令行标准输出
	//os.Stdin.Read()
	fmt.Println("") // os.stdout 的封装
}
