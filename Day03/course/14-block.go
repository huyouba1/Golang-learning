package main

import (
	"fmt"
)

func main() {
	name, desc := "kk", "i'm kk"

	func(name string) {
		fmt.Println(name, desc) // malukang, i'm kk
		name = "烽火"
		desc = "xxx"
		fmt.Println(name, desc) // 烽火，xxx
	}("malukang")

	fmt.Println(name, desc) // kk,xxx

}
