package main

import "fmt"

func main() {
	var names [3]string
	var signIn [3]bool
	var scores [3]float64

	// 类型
	fmt.Printf("%T\n", names)
	fmt.Printf("%T\n", signIn)
	fmt.Printf("%T\n", scores)

	//零值
	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v\n", signIn)
	fmt.Printf("%#v\n", signIn)

	// 字面量
	names = [3]string{"12312312", "我勒个去", "asd2"}
	names = [...]string{"12312312", "我勒个去", "asd2"}
	fmt.Printf("%#v\n", names)

	testnames := [...]string{"12312312", "我勒个去", "asd2"}
	fmt.Printf("%T\n", testnames)

	names = [3]string{1: "huyouba1"} //[3]string{"","huyouba1",""}
	fmt.Printf("%#v\n", names)

	// 操作
	// 关系运算  == !=
	fmt.Println(names == [3]string{})
	fmt.Println(names == [3]string{1: "huyouba1"})

	// 元素
	// 访问 & 修改 索引(0,1,...n-1)
	fmt.Printf("%q\n", names[0])
	names[0] = "02-牛"
	fmt.Printf("%#v\n", names)

	// 函数 len
	fmt.Println(len(names))

	// 遍历
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	fmt.Println("*******************")
	for i, v := range names {
		fmt.Println(i, v)
	}

	// 定义一个数组，他的每一个元素也是一个数组
	// 二维数组
	d2 := [3][2]int{1: [2]int{1, 2}, 0: [2]int{3, 4}, 2: [2]int{1: 5}}
	//[2]int = {0, 0}
	//{[2]int, [2]int, [2]int}
	//{{0, 0}, {0, 0}, {0, 0}}
	fmt.Printf("%#v\n", d2)
	fmt.Printf("%#v\n", d2[1])
	fmt.Printf("%#v\n", d2[1][1])

}
