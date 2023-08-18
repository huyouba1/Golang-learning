package main

import (
	"fmt"
	"regexp"
)

/*
构建regexp结构体
	Compile 编译正则表达式，*Regexp,err
	MustCompile()  *Regexp
*/

func main() {
	reg, err := regexp.Compile("^132\\d{8}$")
	fmt.Println(err, reg)
	// 匹配 Match
	fmt.Println(reg.MatchString("132xxx"))
	fmt.Println(reg.MatchString("12312312312"))
	// 替换 Replace  132??????
	reg, err = regexp.Compile("132\\d{8}")
	fmt.Println(reg.ReplaceAllString("我的电话是132xxx请记录下", "132111111111"))
	fmt.Println(reg.ReplaceAllString("我的电话是13212312312请记录下", "132111111111"))

	// 查找 Find
	fmt.Println(reg.FindAllString("我的电话是13212312312请记录下,13211111111,133,11111111", -1))

	// 分割
	//strings.Split(",")
	reg, err = regexp.Compile("[:;\t,]")
	fmt.Println(reg.Split("我的电话是\t13212312312:请;记录下,13211111111,13311111111", -1))
}
