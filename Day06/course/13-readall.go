package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	file, _ := os.Open("10-multireader.go")
	// 延迟关闭文件
	defer file.Close()

	// 读取文件中的内容
	ctx, _ := ioutil.ReadAll(file)
	fmt.Println(string(ctx))
}
