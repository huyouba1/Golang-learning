package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	//time.Sleep(time.Second * 1)
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("over")
	// 1. 打印0-10
	// 2. 打印全是0  不可能
	// 3. 打印全是10  可能性比较大
	// 4. 随机打印0-9,n个10
}
