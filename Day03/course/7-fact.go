package main

import "fmt"

// 阶乘
// n! =
// n = 0  n! = 1
// n >= 1 n! = n * (n-1)! = n * n-1 * ......* 1
// f(n) = n!
// f(n) = n * f(n - 1)
// 分治  => 大问题分解为多个相同的小问题（小问题可以继续拆分，知道某一个可以解决的子问题）
// 递归调用 => 函数直接或间接调用自己（总有一个停止的条件）

func f(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	fmt.Println(f(3))
}
