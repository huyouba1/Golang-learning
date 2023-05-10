package main

import (
	"fmt"
)

type StructName struct {
	name string
	age  int
}

func (a *StructName) setName(name string) {
	a.name = name
}

func main() {
	aa := &StructName{} // 定义结构体的实例，零值定义{}

	meth1 := (*StructName).setName

	meth1(aa, "asdas")

	fmt.Printf("%#v\n%#v", meth1, aa)
}
