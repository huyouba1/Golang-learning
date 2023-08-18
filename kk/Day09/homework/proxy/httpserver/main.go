package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:9999"
	listenner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listenner.Close()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Printf("[%v] is connected!", conn.RemoteAddr())
		go revReq(conn)
	}

}

func revReq(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	dataCtx := make([]byte, 0)
	for {
		ctx := make([]byte, 1024)
		n, err := reader.Read(ctx)
		if err != io.EOF {
			log.Printf("[%v] has left!\n", conn.RemoteAddr())
			break
		}
		dataCtx = append(dataCtx, ctx[:n]...)
	}
	fmt.Println(string(dataCtx))

}
