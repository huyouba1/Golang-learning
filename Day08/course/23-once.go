package main

import (
	"fmt"
	"sync"
)

func hi() {
	fmt.Println("hi")
}

func main() {
	// 不管函数调用多少次，函数只执行一次
	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(hi)
	}
}
