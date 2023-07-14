package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var locker sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		log.Println("A: Lock Before")
		locker.Lock()

		log.Println("A: Locked")
		time.Sleep(time.Second * 5)

		locker.Unlock()
		log.Println("A: Unlocked")
		wg.Done()
	}()

	go func() {
		log.Println("B: Lock Before")
		locker.Lock()

		log.Println("B: Locked")
		time.Sleep(time.Second * 5)

		locker.Unlock()
		log.Println("B: Unlocked")
		wg.Done()
	}()
	wg.Wait()
}
