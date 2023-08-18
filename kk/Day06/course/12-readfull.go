package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("10-multireader.go")
	defer file.Close()

	// 利用io.Copy + 内存中的流对象
	//buffer := bytes.NewBuffer([]byte(""))

	// strings.Builder
	builder := new(strings.Builder)

	// 复制
	io.Copy(builder, file)

	fmt.Println(builder.String())
}
