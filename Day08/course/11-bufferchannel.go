package main

import (
	"fmt"
	"sync"
)

func main() {
	var channel chan int

	// 带缓冲区的管道 3
	channel = make(chan int, 3)
	fmt.Printf("%T %#v  %d  %d \n", channel, channel, len(channel), cap(channel))

	go func() {
		for i := 0; i < 10; i++ {
			//time.Sleep(time.Second * 2)
			fmt.Println("<-", i)
			channel <- i
		}
		close(channel)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//for i := 0; i < 10; i++ {
		//	time.Sleep(time.Second * 2)
		//	fmt.Println(<-channel)
		//}

		for e := range channel {
			fmt.Println(e)
		}

		wg.Done()
	}()

	wg.Wait()
	//channel <- 1
	//fmt.Printf("%T %#v  %d  %d \n", channel, channel, len(channel), cap(channel))
	//
	//channel <- 1
	//fmt.Printf("%T %#v  %d  %d \n", channel, channel, len(channel), cap(channel))
	//
	//channel <- 1

}
