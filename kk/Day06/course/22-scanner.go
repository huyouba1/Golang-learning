package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个从标准输入读取的 scanner
	scanner := bufio.NewScanner(os.Stdin)

	// 进入循环，直到用户输入 "q"
	for {
		// 提示用户输入内容
		fmt.Print("请输入内容:")

		// 读取用户输入
		scanner.Scan()

		// 如果用户输入 "q"，退出循环
		if "q" == scanner.Text() {
			break
		}

		// 打印用户输入的内容
		fmt.Println("你输入的内容是：", scanner.Text())
	}
}
