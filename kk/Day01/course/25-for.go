package main

import "fmt"

func main() {
	// 控制台打印 1..10
	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}
	//fmt.Println("1")

	// 计算 1..100 的和
	a := 0
	for i := 1; i <= 100; i++ {
		a += i
	}
	fmt.Println(a)
}
