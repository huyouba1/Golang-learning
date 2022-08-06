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
	fmt.Println("********")

	*pointInt = 330000
	fmt.Println(age, age2)
	//fmt.Printf("%T %#v\n", &age, &age)

	pointString = new(string)
	fmt.Printf("%#v,%#v\n", pointString, *pointString)

	pp := &pointString
	fmt.Printf("%T\n", pp)
	**pp = "huyouba1"
	fmt.Println(*pointString)
}
