package main

import (
	"io"
	"os"
)

func main() {
	// 创建文件1
	logFile1, _ := os.Create("test/1.log")
	// 创建文件2
	logFile2, _ := os.Create("test/2.log")
	//logFile, err := os.Create("test/1.log")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 创建multiwriter文件对象，给所有的输出流都写入内容
	writer := io.MultiWriter(logFile1, logFile2, os.Stdout)
	writer.Write([]byte("hello world"))
	//fmt.Println(logFile)
	logFile1.Close()
	logFile2.Close()

}
