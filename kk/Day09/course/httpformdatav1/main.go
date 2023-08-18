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
		// 1. 解析提交内容
		request.ParseMultipartForm(1024 * 1024)

		fmt.Println(request.MultipartForm) // Value, File

		// url 上的数据如何拿取  => Form, FormValue
		// body 值类型 => Form, FormValue, PostForm, PostFormValue
		// file req.MultipartForm.File["name"][0].Open()

		// Todo: 参数检查
		file, _ := request.MultipartForm.File["a"][0].Open()
		io.Copy(os.Stdout, file)
	})
	http.ListenAndServe(":8888", nil)

}
