package main

import (
	"fmt"
	"sort"
)

func main() {
	// 选取切片中第二大元素
	// 1、不去重 选取第二大元素
	// 2、去重 选取第二大元素
	var lst []int = []int{3, 2, 1, 4, 5, 5}
	fmt.Printf("%#v\n", lst)
	sort.Ints(lst)
	fmt.Printf("%#v\n", lst)

	// 第一种，不去重，选取第二大元素
	fmt.Println(lst[len(lst)-1])

}
