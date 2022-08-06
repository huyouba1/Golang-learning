package main

import "fmt"

func main() {
	// 老公

	// 如果有卖西瓜的，买一个包子，否则买十个包子

	var y string
	fmt.Print("有卖西瓜的吗:")
	//y = "yes"
	fmt.Scan(&y)

	if y == "yes" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买十个包子")
	}
}
