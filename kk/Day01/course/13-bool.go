package main

import "fmt"

func main() {
	isGirl := true

	fmt.Printf("%T,%#v", isGirl, isGirl)

	// 操作
	// 逻辑运算

	// 与：左操作数与右操作数都为true的时候则为true  &&
	a, b, c, d := true, true, false, false
	fmt.Println("a,b:", a && b) // true && true = true
	fmt.Println("a,c:", a && c) // true && false = false
	fmt.Println("c,b:", c && b) // false && ture = false
	fmt.Println("c,d:", c && d) // false && false = false
	fmt.Println("************************")

	// 或：左操作数与右操作数有一个为true的时候则为true  ||
	fmt.Println("a,b:", a || b) // true || true = true
	fmt.Println("a,c:", a || c) // true || false = true
	fmt.Println("c,b:", c || b) // false || ture = true
	fmt.Println("c,d:", c || d) // false || false = false
	fmt.Println("************************")

	// 非：取反  ！
	fmt.Println("a:", !a) // true  ! false
	fmt.Println("c:", !c) // false ! true
	fmt.Println("************************")

	// 关系运算
	fmt.Println(a == b) // true == true : true
	fmt.Println(a != c) // true != false : true
	fmt.Println(a == c) // true == false : false
	fmt.Println(c != b) // false != true : true

	fmt.Printf("%t,%t", a, b) //布尔类型使用%t来占位

	var bbbb bool // 布尔类型的零值是false
	println(bbbb)
}
