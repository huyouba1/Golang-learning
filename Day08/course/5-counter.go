// 原子操作函数

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	var locker sync.Mutex

	var count = 5
	var ceil = 10000

	// 10 个例程 5个给counter++  5个给counter--

	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < ceil; i++ {
				locker.Lock()
				counter++
				locker.Unlock()
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
		go func() {
			for i := 0; i < ceil; i++ {
				locker.Lock()
				counter--
				locker.Unlock()
				time.Sleep(time.Microsecond)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
