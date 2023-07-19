package main

import (
	"fmt"
	"sync"
	"worker/pool"
)

// 启动任务
func main() {
	worker := pool.NewPool(2)

	createTask := func(i int) func() interface{} {
		return func() interface{} {
			return i
		}
	}
	for i := 0; i < 5; i++ {
		worker.AddTask(createTask(1))
	}
	//
	//worker.AddTask(func() interface{} {
	//	return 1
	//})
	//
	//worker.AddTask(func() interface{} {
	//	return 2
	//})
	//
	//worker.AddTask(func() interface{} {
	//	return 3
	//})

	worker.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for result := range worker.Results {
			fmt.Println(result)
		}
		wg.Done()
	}()

	worker.Wait()
	wg.Wait()
}
