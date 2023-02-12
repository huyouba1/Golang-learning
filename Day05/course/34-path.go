package main

import "os"

func main() {
	file, err := os.Create("kk.txt")
	file.Close()

	// 二进制程序存放路径 /Users/v_zhenxiyao/Desktop/Golang/Day05/course
	// 相对路径指的是pwd 路径
}
