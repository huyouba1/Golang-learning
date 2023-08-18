package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 重定向  让浏览器重新发起请求到新的地址
	http.HandleFunc("/home/", func(res http.ResponseWriter, req *http.Request) {
		//
		http.Redirect(res, req, "/login/", 302)
		fmt.Fprint(res, "首页")
	})

	http.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "登录页面")
	})

	http.ListenAndServe(":8888", nil)
}
