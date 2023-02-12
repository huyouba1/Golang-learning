package main

import (
	"os"
)

func main() {
	file, err := os.OpenFile("password.txt", os.O_TRUNC|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}

	defer file.Close()

	file.Write([]byte("abc"))
	file.Sync() // 强制落盘
	file.Write([]byte("123"))
	file.WriteString("abcdefghijklmn")
}
