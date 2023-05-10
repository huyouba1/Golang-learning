package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 通过New方法创建结构体指针对象
	reader := bytes.NewReader([]byte("abcdef"))
	ctx := make([]byte, 3)
	n, _ := reader.Read(ctx)

	fmt.Println(string(ctx[:n]))

	// 重置reader中的内容
	reader.Reset([]byte("123456"))

	// 输出内容到输出流
	reader.WriteTo(os.Stdout)
}
