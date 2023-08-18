package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 4, 5, 1, 2, 8, 9}
	//sort.Sort(sort.Reverse(sort.IntSlice(nums)))   // 倒序
	//fmt.Println(nums)

	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"x", "t", "z", "q", "a", "s"}
	sort.Strings(names)
	fmt.Println(names)

	// 查找
	// [0,100)
	// 50  大[0,50)之间
	// 24  小[24,50]之间
	// x
	// 0,1,2,3,4,5.....100

	// 二分查找
	// [1, 3, 5, 9, 10] // 有序 0...4 => 2
	// x 是不是在切片中 x = 5
	// x = 3 [0, 4] 2 [0, 2) 0 (0, 2) 1 => 找到
	// x = 8 [0, 4] 2 (2, 4] 3 (2, 3) 没找到
	nums = []int{1, 3, 5, 9, 10}
	fmt.Println(nums[sort.SearchInts(nums, 8)] == 8)
	fmt.Println(nums[sort.SearchInts(nums, 5)] == 5)

}
