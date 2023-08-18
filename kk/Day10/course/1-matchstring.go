package main

import "regexp"
import "fmt"

func main() {
	// . 任意
	// \d 数字  \D 非数字
	// \w 数字，大小写英文字母 _
	// \S 空白字符  \s 非空白字符
	// \d 数字 0,1,2,3,4,5,6,7,8,9    0|1|2|3|4|5|6|7|8|9  [0123456789][0-9]
	// [a-z]
	// [^a-z] 取反
	// ?  0个或1个
	// +  至少1个
	// *  任意多个
	// {n,m}  字符数量≥n ≤m
	var pattern string = "^132" // 以132开头
	fmt.Println(regexp.MatchString(pattern, "132zzzzz"))
	fmt.Println(regexp.MatchString(pattern, "1222zzzzz"))

	// 以132开头，都是数字，长度为11位 [0-9]
	pattern = "^132\\d{8}$"
	fmt.Println(regexp.MatchString(pattern, "132zzzzz"))
	fmt.Println(regexp.MatchString(pattern, "132123123"))
	fmt.Println(regexp.MatchString(pattern, "13212331233"))
	fmt.Println(regexp.MatchString(pattern, "13212331233x"))

	fmt.Println("==============")
	// 132 158
	// 1[35][28] 132 138 152 158
	// 分组 ()
	// ^(132)|(158)[0-9]{8}$
	pattern = "^132|^158\\d{8}$"
	fmt.Println(regexp.MatchString(pattern, "13212331233"))
	fmt.Println(regexp.MatchString(pattern, "15812331233"))
	fmt.Println(regexp.MatchString(pattern, "13212331233x"))
	// 邮箱格式
	// xxxx@xxx.xx
	// xxxx(@之前) 数字，大小写英文字母长度1-32字符
	// xxx(.之前) 小写英文字母组成 长度1-12字符
	// xx（.之后） edu
	fmt.Println("=======email=======")
	//pattern = "^[A-Za-z0-9]{1,32}@[a-z]{1,12}\\.edu$"
	pattern = "^[a-zA-Z0-9]{1,32}@[a-z]{1,12}[.]edu$"

	fmt.Println(regexp.MatchString(pattern, "aasdasf@baidu.edu"))
	fmt.Println(regexp.MatchString(pattern, "asdalskfj1231@1.edu"))
	fmt.Println(regexp.MatchString(pattern, "?@baidu.edu"))
	fmt.Println(regexp.MatchString(pattern, "我才asdalskfj1231@baidux.edua"))

	fmt.Println(regexp.QuoteMeta(pattern))
}
