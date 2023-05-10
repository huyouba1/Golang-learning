package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ctx, err := ioutil.ReadFile("10-multireader.go")
	fmt.Println(string(ctx), err)
}
