package main

import (
	"log"
	"os"
)

func main() {
	// 设置格式的
	// flags
	log.SetFlags(log.Flags() | log.Ldate | log.Lshortfile) // 全局日志配置
	// prefix
	log.SetPrefix("main: ")

	log.Println("我是第一条Println日志")

	//log.Fatalln("我是一个Fatal日志")   // 打印日志之后，直接退出
	//log.Panic("我是一条panic日志") // 发生了错误，但是可以恢复，不退出应用

	log.Println("我是第二条Println日志")

	// DEBUG,INFO,WARNING,ERROR
	// logrus

	logger := log.New(os.Stdout, "logger: ", log.Flags())
	logger2 := log.New(os.Stdout, "logger2: ", log.Flags())

	//logger.SetPrefix("logger: ")
	//logger2.SetPrefix("logger2: ")

	logger.Println("我是logger日志")
	logger2.Println("我是logger2日志")

	// 标准输入、输出  fmt.Scan 输入    fmt.Println  输出
	// os.Stdin, os.Stdout,os.Stderr
}
