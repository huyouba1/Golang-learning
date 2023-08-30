package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"os"
)

type Request struct {
	A   int
	B   int
	Sum int
}

func main() {
	dialer := &websocket.Dialer{}
	haeder := http.Header{
		"Name": []string{"zzz"},
	}
	conn, resp, err := dialer.Dial("ws://localhost:5657", haeder)
	if err != nil {
		fmt.Println("建立连接失败", err)
		fmt.Println(resp.StatusCode)
		io.Copy(os.Stdout, resp.Body)
		return
	}
	for k, v := range resp.Header {
		fmt.Printf("ket == %s, value == %s\n", k, v)
	}

	defer conn.Close()
	for i := 0; i < 5; i++ {
		request := Request{A: 3, B: 5}
		conn.WriteJSON(request)
		var response Request
		conn.ReadJSON(&response)
		fmt.Println(response.Sum)
	}

}
