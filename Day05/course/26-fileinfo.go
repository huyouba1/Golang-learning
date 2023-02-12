package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, _ := os.Stat("password")
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.ModTime())
}
