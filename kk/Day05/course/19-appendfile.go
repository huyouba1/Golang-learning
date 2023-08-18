package main

import (
	"fmt"
	"os"
)

// os.O_CREATE  没有则创建
// os.O_APPEND  追加
// os.O_TRUNC   覆盖
// os.O_RDWR    读写

func main() {
	file, err := os.OpenFile("name.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write([]byte("abcd"))
	file.Write([]byte("abcd"))
	file.Write([]byte("abcd"))
	file.Write([]byte("abcd"))

}
