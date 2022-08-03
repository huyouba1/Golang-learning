package main

import "fmt"

func main() {
	var (
		// 零值 nil
		pointInt    *int
		pointString *string
	)

	fmt.Printf("%T %#v\n", pointInt, pointInt)
	fmt.Printf("%T %#v\n", pointString, pointString)

	// 赋值
	// 使用现有变量，取地址 &name
	age := 32
	age2 := age
	fmt.Printf("%T %#v \n", &age, &age)
	fmt.Printf("%T %#v \n", &age2, &age2)

	pointInt = &age
	fmt.Println(pointInt)
	fmt.Println(*pointInt)

	*pointInt = 33
	fmt.Println(age, age2)
}
