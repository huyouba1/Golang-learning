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

	// 有编码格式
	// application/x-www-form-urlencoded

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// 1. 解析参数
		request.ParseForm()
		// 2. 获取
		fmt.Println(request.Form) // 包含请求体和URL中的数据
		fmt.Println(request.Form.Get("a"))
		fmt.Println(request.Form["a"])
		fmt.Println(request.PostForm) // 只包含请求体中数据
	})
	http.ListenAndServe(":8888", nil)

}
