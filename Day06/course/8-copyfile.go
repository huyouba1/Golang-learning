package main

import (
	"fmt"
	"io"
	"os"
)

// 复制文件(src -> dest)
func CopyFile(src, dest string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// 延迟关闭
	defer srcFile.Close()

	// 以写方式打开目的文件
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	// 延迟关闭
	defer destFile.Close()

	// 这两个是copy所有
	// 从srcFile赋值内容到destFile
	_, err = io.Copy(destFile, srcFile)
	//buffer := make([]byte, 1024*1024)
	//_, err = io.CopyBuffer(destFile, srcFile, buffer)

	// 只copy前N个字节
	//_, err = io.CopyN(destFile, srcFile, 1024*1024)

	return err
}

func main() {
	fmt.Println(CopyFile("os.go", "os.go5"))

}
