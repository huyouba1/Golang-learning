package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, _ := os.Open("10-multireader.go")
	//延迟关闭
	defer file.Close()

	// 定义buffer，用于每次读取文件内容
	buffer := make([]byte, 10)

	// 用于存储文件内容
	ctx := make([]byte, 0, 1024*1024)

	// 分批读取文件内容放入ctx 中
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		ctx = append(ctx, buffer[:n]...)
		//fmt.Println(buffer[:n])
	}
	fmt.Println(string(ctx))
	fmt.Println()
}
