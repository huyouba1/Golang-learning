package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 种子只需要设置一次
	rand.Seed(time.Now().Unix())

	// [0-100]
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int() % 101) // 取模
	}

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
