package main

import (
	"fmt"
	"sync"
)

func main() {
	var smap sync.Map
	smap.Store("1", 1)
	fmt.Println(smap.Load("1"))
	smap.Delete("1")
	fmt.Println(smap.Load("1"))

	smap.Store("2", 2)
	smap.Store("3", 3)
	smap.Store("4", 4)
	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
