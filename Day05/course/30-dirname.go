package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("password")
	if err != nil {
		return
	}

	defer file.Close()
	names, err := file.Readdirnames(-1) // 字符串的切片
	fmt.Println(names, err)
}
