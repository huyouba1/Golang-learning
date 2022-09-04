package main

import "fmt"

// 赋值 <==> 实参/形参
func test1(n int) {

	n = 1
}

func test2(s []int) {
	fmt.Printf("%p\n", s)
	s[0] = 1
}

func main() {
	a := 0
	b := make([]int, 10)
	test1(a)
	test2(b)
	fmt.Println(a)
	fmt.Printf("%p\n", b)
	fmt.Println(b)
	// 值类型  b = a(在内存中申请新的内存空间，将a的值copy到b中)
	// 在修改a的时候，不影响b
	// 在修改a的时候，影响b

	// 引用类型  b = a (赋值的地址)
	// b => a 存储内容
	//age := 30
	//temAge := age
	//
	//temAge = 31
	//fmt.Println(age, temAge)
	//
	//users := make([]string, 10)
	//temUsers := users
	//temUsers[0] = "kk"
	//fmt.Printf("%#v\n%#v\n", users, temUsers)
	//
	//// 值类型
	//// int, float, point, 数组, 结构体类型
	//
	//// 引用类型
	//// 切片, 映射, 接口
	//array := [3]int{}
	//tmpArray := array
	//tmpArray[0] = 10
	//fmt.Println(array, tmpArray)

}
