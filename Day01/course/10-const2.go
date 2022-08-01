package main

import "fmt"

func main() {

	const (
		A = "test"
		B // 如何省略赋值，则会使用前一个常量的值进行初始化(A)
		C // 如何省略赋值，则会使用前一个常量的值进行初始化(B=A)
		D = "testD"
		E // 如何省略赋值，则会使用前一个常量的值进行初始化(D)
		F // 如何省略赋值，则会使用前一个常量的值进行初始化(E=D)
	)
	fmt.Println(A, B, C, D, E, F)
}
