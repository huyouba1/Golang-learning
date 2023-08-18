package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int int8  type()
	// float32 float64
	//string
	var (
		intVal     = 65
		float64Val = 2.2
		stringVal  = "3.3"
	)
	fmt.Println(intVal, float64Val, stringVal)
	fmt.Printf("%T %#v\n", float64(intVal), float64(intVal))
	fmt.Printf("%T %#v", int(float64Val), int(float64Val))

	fmt.Println("******")
	fmt.Println(string(intVal))
	//fmt.Println(int(stringVal))   // 报错
	//fmt.Println(float64(stringVal)) // 报错
	v, err := strconv.Atoi(stringVal) // 字符串转换为int
	fmt.Println(err, v)

	vv, err := strconv.ParseFloat(stringVal, 64) // 字符串转换成float64
	fmt.Println(err, vv)

	fmt.Println(strconv.Itoa(intVal))                         // 把int转换成字符串
	fmt.Println(strconv.FormatFloat(float64Val, 'f', 10, 64)) // 把float 转换为字符串

}
