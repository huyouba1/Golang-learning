package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	fmt.Println(runtime.GOMAXPROCS(1)) // 限制cpu使用
	fmt.Println(runtime.GOMAXPROCS(1))
}
