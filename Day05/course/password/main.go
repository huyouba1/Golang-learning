package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"password/uitls"
)

const hash = "YQvjgh$$caf6c7eadbbc55f0316ef061c3988380"

func main() {
	fmt.Print("请输入密码: ")
	password, _ := gopass.GetPasswd()
	if utils.CheckPassword(string(password), hash) {
		fmt.Println("成功")
	} else {
		fmt.Println("失败")
	}
}
