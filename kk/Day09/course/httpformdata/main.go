package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// http.post
	// 提交数据 请求体中
	// 有编码格式
	// application/x-www-form-urlencoded
	// k=v&k2=v2
	// 上传文件 => multipart/form-data
	// application/json {"a" : 1}

	// 上传文件 => multipart/form-data

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		file, header, _ := request.FormFile("a")

		io.Copy(os.Stdout, file)
		fmt.Println(header.Filename)
		fmt.Println(header.Size)
		fmt.Println(header.Header)
	})
	http.ListenAndServe(":8888", nil)

}
