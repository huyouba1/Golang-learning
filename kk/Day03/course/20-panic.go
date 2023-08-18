package main

import (
	"fmt"
)

func test() (err error) {
	//recover 必须在延迟函数内
	defer func() {
		fmt.Println("defer")
		if panicErr := recover(); panicErr != nil {
			err = fmt.Errorf("%s", panicErr)
		}
	}()
	fmt.Println("before")
	//panic("自定义panic")
	//err := errors.New("xxx")
	// 检查所有的数据，连接，资源。。。
	// 用别人的库的时候，panic
	fmt.Println("after")
	//return err
	return

}

func main() {
	fmt.Println("befor main")
	err := test()
	fmt.Println("after main", err)
}
