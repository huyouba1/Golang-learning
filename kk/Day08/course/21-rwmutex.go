package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// 写 读
	// 读 读
	// 在并发的什么场景下会出现数据混乱  = 改

	// 读写锁
	var locker sync.RWMutex
	// 获取锁 Lock.Rlock
	// 释放锁 Unlock,RUnlock

	// 写 读  Rlock Lock 也是互斥的
	// 写 写  Lock Lock  也是互斥的
	// 读 读 Rlock Rlock

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		log.Println("A: Lock Before")
		locker.RLock()

		log.Println("A: Locked")
		time.Sleep(time.Second * 5)

		locker.RUnlock()
		log.Println("A: Unlocked")
		wg.Done()
	}()

	go func() {
		log.Println("B: Lock Before")
		locker.RLock()

		log.Println("B: Locked")
		time.Sleep(time.Second * 5)

		locker.RUnlock()
		log.Println("B: Unlocked")
		wg.Done()
	}()
	wg.Wait()
}
