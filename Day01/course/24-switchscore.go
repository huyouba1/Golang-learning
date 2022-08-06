package main

import "fmt"

func main() {
	// 在控制台输入分数
	// 90 以上 A
	// 80 以上 B
	// 70 以上 C
	// 60 以上 D
	var score float32
	fmt.Print("请输入一个分数:")
	fmt.Scan(&score)
	fmt.Println("你输入的分数是:", score)

	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("C")
	default:
		fmt.Println("D")

	}
	// if 必须有
	// else if 可以有0-N
	// else
	//if score >= 90 {
	//	fmt.Println("A")
	//} else if score >= 80 {
	//	fmt.Println("B")
	//} else if score >= 60 {
	//	fmt.Println("C")
	//} else {
	//	fmt.Println("D")
	//}
}
