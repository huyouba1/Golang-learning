package main

import (
	"encoding/json"
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

	// 上传文件 => multipart/form-data
	// application/json {"a" : 1}

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// 需要读取请求体中数据
		//io.Copy(os.Stdout, request.Body)
		decoder := json.NewDecoder(request.Body)
		var info map[string]interface{}
		{
		}
		decoder.Decode(&info)
		fmt.Println(info)
	})
	http.ListenAndServe(":8888", nil)

}
