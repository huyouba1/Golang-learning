package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 当请求URL未绑定，按照URL中最近匹配的绑定关系去处理
	/*
			/  indexHandleFunc
		    /time/ timeHandleFunc

			/time/ => timeHandleFunc
			/time/xxxx/xxxx =>  /time/ => timeHandleFunc

			/abc/abc/ => / => indexHandleFunc
	*/
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL, request.Proto)

		fmt.Printf("%T,%#v\n", request.Header, request.Header)
		fmt.Println(request.Header.Get("User-Agent"))

	})
	http.ListenAndServe(":8888", nil)
}
