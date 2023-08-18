package main

import "fmt"

const PackageVar string = "packageVar"
const PackMsg = "msg"

func main() {
	const name string = "huyouba1"
	const msg = "msg"
	fmt.Println(name)

	name = "selicen" // 常量是不能修改的，只能在初始化的时候赋值

}
