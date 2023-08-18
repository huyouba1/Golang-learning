package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 5)
		// 关闭管道，不再写入
		close(channel)
		//channel <- struct{}{}
	}()
	fmt.Println("before")
	_, ok := <-channel // 当管道关闭，读取管道返回数据，不再进行阻塞
	fmt.Println("over")
	fmt.Println(ok)
	time.Sleep(time.Second * 5)
}
