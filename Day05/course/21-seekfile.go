package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("password.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//fmt.Println(file.Seek(0, os.SEEK_CUR))
	file.Seek(3, os.SEEK_SET) // 设置游标偏移量
	ctx := make([]byte, 3)
	n, err := file.Read(ctx)
	fmt.Println(n, err, string(ctx))

	file.Seek(0, os.SEEK_SET)
	n, err = file.Read(ctx)
	fmt.Println(n, err, string(ctx))
	fmt.Println(file.Seek(0, os.SEEK_CUR)) // 打印当前游标位置
	file.Write([]byte("abc"))
}
