package main

import "fmt"

func main() {
BREAK:
	for j := 0; j <= 3; j++ {
		fmt.Println("__________")
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
			if i == 3 {
				break BREAK
			}
		}
	}
}
