package main

import "fmt"

func main() {
	/*
			标识符：程序中定义的名字、变量名、常量名、函数名字、自定义类型
			规范：
				1. 必须满足：组成只能由非空的Unicode编码字符串、数字、下划线组成
				2. 必须满足：必须以Unicode编码的字符串或下划线开头（不能以数字开头）
				3. 必须满足：不能以go的关键字冲突（package，func，var。。。25个）
			建议：
				1. ASCILL编码（a-z，A-Z）、数字、下划线
				2. 变量使用驼峰式，多个英文字母 myname myName my_name
				3. 与go内置标识符不要冲突(string)
			说明：标识符区分大小写
				my = ""
		        MY = ""
	*/
	var my = "my"
	var My = "My"

	var 我的名字 = "huyouba1"    //可以
	var _myname = "huyouba1" //可以
	//var 0name = "huyouba1"        //不可以
	//var package = "huyouba1" //不可以
	fmt.Println(my, My)
	fmt.Println(我的名字, _myname)
}
