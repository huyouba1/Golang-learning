package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Executable())
	fmt.Println(os.Getwd())
}
