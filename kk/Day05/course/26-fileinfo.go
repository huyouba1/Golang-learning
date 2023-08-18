package main

import (
	"fmt"
	"os"
)

func main() {
	fileinfo, _ := os.Stat("password.txt")
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Size())
	fmt.Println(fileinfo.ModTime())
}
