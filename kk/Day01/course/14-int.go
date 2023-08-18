package main

import "fmt"

func main() {
	var age8 int8 = 18 //字面量 10 进制   8进制   16进制
	var age int = 18
	fmt.Printf("%T,%#v,%d\n", age8, age8, age8)
	fmt.Printf("%T,%#v,%d\n", age, age, age)

	/*
			了解：
			8 进制： 0 X X X < 8
			16 进制： 0 X X X  < 16 0-9 A=10 B=11 C=12 D=13
			070 => 10 进制:
			078 => 10 进制:

		fmt.Println(070, 78)

		// 二进制存储
		base 2
		7 = 0111
	*/

	// 常用操作
	// 算数运算 + - * / % ++ --
	a, b := 2, 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 0
	fmt.Println(a % b) // 2

	a++
	b--
	fmt.Println(a, b) // 3 3

	// 关系运算 >  < >=  <= == !=
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // false
	fmt.Println(a <= b) // true
	fmt.Println(a >= b) // true
	fmt.Println(a == b) // true
	fmt.Println(a != b) // false

	/*
		仅了解 位运算
		7 = 0111
		2 = 0010

		复数：二进制表示补码，对应正数，取反+1
		-3 3 =》0011 =》 1101

		按位于：&,两个都为1，结果为1
		按位或：|，有一个为1则为1
		取非： ^，对于有符号首字符不变，1->0,0->1
		右移位: >>  7 >> 2 0001 => 1
		左移位: <<
		and not: &^
	*/
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(^2)
	fmt.Println(7 >> 2)
	fmt.Println(7 << 2)
	fmt.Println(7 &^ 2)

	//fmt.Println(age + age8)   // 不同的数据类型是不能进行计算的
	var (
		i   int   = 1
		i32 int32 = 1
		i64 int64 = 1
	)
	// 类型转换 type(value)  int32(i) int(i32)
	fmt.Println(i + int(i32) + int(i64))
	fmt.Println(int32(i) + i32 + int32(i64))
	fmt.Printf("%T\n", int32(i)+i32+int32(i64))
	fmt.Printf("%T\n", i+int(i32)+int(i64))

	var (
		achar        byte = 'A'
		aint         byte = 64
		unicodePoint rune = '中'

		// 字符串 =》内存(0101) =》 转换 =》 编码(utf8,utf16,gbk,gb2312)
	)
	fmt.Println(achar, aint, unicodePoint)
	fmt.Printf("%d, %b, %o ,%x ,%U , %c， %c", achar, 15, 15, 15, unicodePoint, achar, 65)
}
