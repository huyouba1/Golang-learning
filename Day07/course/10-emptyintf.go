package main

import (
	"fmt"
)

// 空接口,不包含任何函数签名
type EmptyIntf interface {
}

func PrintType(v interface{}) {
	switch value := v.(type) {
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("boot", value)
	case [3]int:
		fmt.Println("[3]int", value)
	case []int:
		fmt.Println("[]int", value)
	case map[string]string:
		fmt.Println("map[string]string", value)
	default:
		fmt.Println("unknow", value)
	}
}

func main() {
	var emptyIntf EmptyIntf
	emptyIntf = 1
	emptyIntf = true
	emptyIntf = "aa"
	fmt.Println(emptyIntf)

	var emptyIntf2 interface{}
	emptyIntf2 = 1
	emptyIntf2 = "test"
	fmt.Println(emptyIntf2)

	PrintType(1)
	PrintType(false)
	PrintType([]int{3, 2, 1})
	PrintType(map[string]string{"abcd": "acaa"})
}
