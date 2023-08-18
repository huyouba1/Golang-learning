package main

import "fmt"

func main() {
	var msg = "我的名字是\\n huyouba1"
	var msgRaw = `我的名字是\n\n huyouba1`
	fmt.Printf("%T %s \n", msg, msg)
	fmt.Printf("%T %s \n", msgRaw, msgRaw)

	// 操作
	// 字符串连接 +
	fmt.Println(msg + msgRaw)

	// 关系运算 > >= = <= <  != ==
	fmt.Println("abc" > "acd") // false

	// 赋值 = += -=
	msg += "----abcd3"
	fmt.Println(msg)

	// 索引  切片 必须是ASCII码
	msg = "abcdef"
	fmt.Printf("%T %#v %c", msg[0], msg[0], msg[0])
	fmt.Println("*********")
	fmt.Println(msg[1:3])

	// len  字节的大小,并不是字符数量
	fmt.Println(len(msg))
	fmt.Println(len(msgRaw))
}
