package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test")
	fmt.Println(os.MkdirAll("test/a/c", os.ModePerm))
	fmt.Println(os.Remove("test/a/c"))
	os.RemoveAll("test")
}
