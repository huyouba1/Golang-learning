package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-1))     // 获取绝对值
	fmt.Println(math.Ceil(1.5))   // 向上取整
	fmt.Println(math.Ceil(-1.5))  // 向上取整
	fmt.Println(math.Floor(1.5))  // 向下取整
	fmt.Println(math.Floor(-1.5)) // 向下取整
	fmt.Println(math.Max(1, 2))   // 取最大值
	fmt.Println(math.Min(1, 2))   // 取最小值
}
