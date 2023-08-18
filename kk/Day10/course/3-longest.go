package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg, _ := regexp.Compile("[ab0-9]+")
	fmt.Println(reg.FindAllString("0-a23-b3456-a23b3456", -1)) // 贪婪模式 [0 a23 b3456 a23b3456]

	reg, _ = regexp.Compile("(?U)[ab0-9]+") //?U   [0 a 2 3 b 3 4 5 6 a 2 3 b 3 4 5 6]
	fmt.Println(reg.FindAllString("0-a23-b3456-a23b3456", -1))

	// 将非贪婪模式转换为贪婪模式
	reg.Longest()
	fmt.Println(reg.FindAllString("0-a23-b3456-a23b3456", -1))
}
