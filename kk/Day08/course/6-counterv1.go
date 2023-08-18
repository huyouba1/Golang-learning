// 原子操作函数

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var counter int64

	var count = 5
	var ceil = 10000

	// 10 个例程 5个给counter++  5个给counter--

	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < ceil; i++ {
				atomic.AddInt64(&counter, 1)
				//counter++
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < ceil; i++ {
				//counter--
				atomic.AddInt64(&counter, -1)
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
