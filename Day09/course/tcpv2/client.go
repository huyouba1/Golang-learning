package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const dateFormat = "2006-01-02 15:04:05"

func main() {
	addr := "127.0.0.1:9999"

	// 连接服务器
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 给服务器发送数据
	log.Printf("connected")
	fmt.Fprintf(conn, "Time1: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time2: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintf(conn, "Time3: %s\n", time.Now().Format(dateFormat))
	fmt.Fprintln(conn, "quit\n")

	// 读取服务器端数据
	reader := bufio.NewReader(conn)
	line, _, err := reader.ReadLine()
	fmt.Println(string(line))
}
