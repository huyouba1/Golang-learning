package main

import (
	"io"
	"net/http"
	"time"
)

// 1.定义处理器 必须满足HandlerFunc interface 自定义类型（结构体）
// Handler接口
// ServerHTTP(http.ResponseWriter, *http.Request)

type TimeHandler struct {
}

func (h *TimeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	io.WriteString(response, now)
}

func main() {
	http.Handle("/time/", &TimeHandler{})

	http.ListenAndServe(":9998", nil)
}
