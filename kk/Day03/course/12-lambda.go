package main

import "fmt"

func calc(n1 int, n2 int, callback func(int, int) int) int {
	// 不定义是什么运算
	// 通过函数参数传递我要进行的运算
	rt := callback(n1, n2)
	// 检查结果在 0 -- 100 范围内，如果超过返回 -1
	if rt >= 0 && rt <= 100 {
		return rt
	}
	return -1
}

func main() {
	rt := calc(50, 30, func(n1, n2 int) int { // 匿名函数
		return n1 + n2
	})
	fmt.Println(rt)

	rt = calc(50, 30, func(n1, n2 int) int {
		return n1 * n2
	})
	fmt.Println(rt)
}
