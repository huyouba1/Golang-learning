package main

import (
	"fmt"
	"time"
)

func timeTick(interval time.Duration) <-chan time.Time {
	timeChannel := make(chan time.Time)
	go func() {
		for {
			time.Sleep(interval)
			timeChannel <- time.Now()
		}

	}()
	return timeChannel
}

func main() {
	endTime := time.Now().Add(time.Second * 30)
	for now := range timeTick(time.Second * 3) {
		fmt.Println("now:", now)
		//fmt.Println("endTime:", endTime)
		//if time.Now().After(endTime) {
		//	break
		//}
		if now.After(endTime) {
			break
		}
	}
}
