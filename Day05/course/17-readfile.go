package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 文件路径
	// 相对路径  程序的执行路径
	// password
	// ./password
	// ../password
	// 绝对路径  根目录、磁盘目录写起的路径
	// /opt/todolist/etc/password
	// e:\todolist\etc\password

	// 读文件
	// 程序 不会停止退出，重复的读取文件

	file, err := os.Open("password.txt")
	if err != nil {
		return
	}
	defer file.Close()

	cxt := make([]byte, 1024)
	for {
		//fmt.Println(file, err)

		n, err := file.Read(cxt)
		if err == io.EOF {
			break
		}
		fmt.Println(n, err, string(cxt[:n]), len(string(cxt[:n])))
	}
}
