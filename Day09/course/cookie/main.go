package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func parseCookie(cookie string) map[string]string {
	cookieMap := make(map[string]string)
	if strings.TrimSpace(cookie) == "" {
		return cookieMap
	}

	values := strings.Split(cookie, ";")
	// Todo: 格式检查
	for _, value := range values {
		es := strings.Split(value, "=")
		cookieMap[strings.TrimSpace(es[0])] = strings.TrimSpace(es[1])
	}
	return cookieMap
}

func main() {
	// cookie 浏览器端存储
	// 读取cookie中数据 counter
	// counter += 1 设置在浏览器中(counter无,设置为0+1)

	// 请求获取： Cookie
	// 响应设置： Set-Cookie

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// 读取cookie
		cookie := parseCookie(req.Header.Get("Cookie"))
		counter := 0
		if v, err := strconv.Atoi(cookie["counter"]); err == nil {
			counter = v
		}
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter),
			HttpOnly: true,
		}
		http.SetCookie(res, counterCookie)
		fmt.Fprintf(res, "counter: %d", counter)
	})

	http.ListenAndServe(":8888", nil)
}
