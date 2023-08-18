package main

import "fmt"

// 无参  无返回值
func sayHello() {
	fmt.Println("Hello World")

}

// 有参数 无返回值
func sayHi(name string, name2 string) {
	fmt.Println("Hi:", name, name2)

}

// 有参数 有返回值
func add(n1 int, n2 int) int {
	return n1 + n2

}

//
func test(a int, b string) {
	fmt.Println(a, b)

}

func main() {
	//sayHi() // 函数调用
	a := sayHi // 赋值
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", sayHi)
	a("1", "2")

	var callback func(int, int) int
	fmt.Printf("%T,%#v\n", callback, callback)

	fmt.Printf("%T\n", sayHello)

	callback = add
	fmt.Printf("%T,%#v\n", callback, callback)
	rt := callback(1, 4)
	fmt.Println(rt)
}
