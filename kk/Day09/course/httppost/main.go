package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.post
	// 提交数据 请求体中
	// 有编码格式
	// application/x-www-form-urlencoded
	// k=v&k2=v2
	// 上传文件 => multipart/form-data
	// application/json {"a" : 1}

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request.PostFormValue("a"))
	})
	http.ListenAndServe(":8888", nil)

}
