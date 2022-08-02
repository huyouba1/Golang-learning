package main

import "fmt"

func main() {
	var name = "huyouba1"
	fmt.Println(name)
	fmt.Println("*") // 打印变量+换行
	fmt.Print(name)
	fmt.Print("*") // 打印变量不换行
	fmt.Println("*")
	fmt.Printf("%T,%v,%#v", name, name, name)
}
