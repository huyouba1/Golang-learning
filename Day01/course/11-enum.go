package main

import "fmt"

func main() {
	const (
		Sun = iota // 在常量组内使用，iota初始化为0，每次调用的时候+1
		Mon
		Tuesd
		Wed
		Thur
		Fri
		Sta
		aa = iota //依旧会累计
	)
	fmt.Println(Sun, Mon, Tuesd, Wed, Thur, Fri, Sta, aa)
}
