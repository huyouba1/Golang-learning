package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Request struct {
	A   int
	B   int
	Sum int
}

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
	}
	// 没给listener赋值
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil) // 将http协议升级到websocket协议
	if err != nil {
		fmt.Printf("升级失败 %s\n", err)
		return
	}
	fmt.Printf("跟客户端 %s  建立好了websocket连接", r.RemoteAddr)
	go ws.handleOneConnection(conn)
}

func (ws *WsServer) handleOneConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()
	for {
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		var request Request
		if err := conn.ReadJSON(&request); err != nil {
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Println("发生了都超市")
					return
				}
			}
			fmt.Println(err)
			return
		}
		response := Request{Sum: request.A + request.B}
		err := conn.WriteJSON(response)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (ws *WsServer) Start() error {
	ws.listener, _ = net.Listen("tcp", ws.addr)
	err := http.Serve(ws.listener, ws)

	//err := http.ListenAndServe(ws.addr, ws)
	if err != nil {
		fmt.Println("server failed")
		return err
	}
	return nil
}

func main() {
	ws := NewWsServer(5657)
	ws.Start()
}
