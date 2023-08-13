package main

import (
	"fmt"
	"net/http"
)

func main() {

	// http 协议
	// GET
	// url?params
	// params k=>v&k2=v2

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// 1. 解析参数
		request.ParseForm()
		// 2. 获取
		fmt.Println(request.Form)
		fmt.Println(request.Form.Get("a"))
		fmt.Println(request.Form["a"])
	})
	http.ListenAndServe(":8888", nil)
}
