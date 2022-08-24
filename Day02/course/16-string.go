package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	ascii := "abc我爱中华人民共和国"
	fmt.Println([]byte(ascii))
	fmt.Println([]rune(ascii))

	fmt.Println(len(ascii))
	fmt.Println(len([]rune(ascii)))            // 统计Unicode编码格式的数量
	fmt.Println(utf8.RuneCountInString(ascii)) // 统计Unicode编码格式的数量

	// byte/rune 转换为字符串
	fmt.Println(string([]byte(ascii)))
	fmt.Println(string([]rune(ascii)))

	// int,float,bool
	fmt.Println(strconv.Itoa('a')) // 数字转字符
	ch, err := strconv.Atoi("97")  // 字符串转 Int
	fmt.Println(ch, err)

	fmt.Println(strconv.FormatFloat(3.1415926, 'f', 10, 64)) // 浮点数转字符串
	pi, err := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(pi, err) // 字符串转浮点数

	fmt.Println(strconv.FormatBool(true)) // 布尔转字符
	b, err := strconv.ParseBool("true")   // 字符转布尔
	fmt.Println(b, err)

	fmt.Println(strconv.FormatInt(15, 2))  // Int 转2进制
	fmt.Println(strconv.FormatInt(15, 16)) // Int 转16进制

}
