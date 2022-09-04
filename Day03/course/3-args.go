package main

import "fmt"

// 可变参数的函数

func test1(args ...string) {
	fmt.Printf("%T,%#v\n", args, args)
	//fmt.Println(args)
}

// 1. 可变参数在一个方法中，只能有一个，并且可变参数必须放在函数声明参数列表最后

// 举例(函数至少有n个参数)
// add(n1,n2,....)
func add1(n1, n2 int, args ...int) int {
	total := n1 + n2
	for _, v := range args {
		total += v
	}
	return total
}

func calc(n1, n2 int, args ...int) int {
	// 直接调用add1方法，把add1方法结果进行返回
	// args 切片
	// add(n1,n2,args[0],args[1],...)
	return add1(n1, n2, args...) // 解操作
}

func main() {
	test1()
	test1("1")
	test1("1", "2", "3")

	fmt.Println(add1(1, 2))
	fmt.Println(add1(1, 2, 3, 4))
	fmt.Println(add1(1, 2, 3, 4, 10))
	fmt.Println(calc(1, 2, 3, 4))

	params := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(add1(1, 2, params...))
}
