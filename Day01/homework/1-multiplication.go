package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for x := 1; x <= i; x++ {
			fmt.Printf("%d*%d=%-4d", x, i, x*i)
		}
		fmt.Println()
	}
}
