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

	// Formvalue 只能获取key对应的一个值

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request.FormValue("a"))
	})
	http.ListenAndServe(":8888", nil)
}
