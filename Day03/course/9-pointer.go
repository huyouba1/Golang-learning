package main

import "fmt"

// 值类型 在函数内修改实参
// fmt.Scan(&age)

func change(value int) {
	value += 1

}

func changePointer(point *int) {
	*point = *point + 1
}
func main() {
	value := 1
	change(value)
	fmt.Println(value) // 1

	changePointer(&value)
	fmt.Println(value)
}
