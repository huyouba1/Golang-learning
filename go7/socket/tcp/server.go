package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5656")
	checkError(err)
	listen, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	conn, err := listen.Accept()
	checkError(err)

	request := make([]byte, 256)
	n, err := conn.Read(request)
	checkError(err)
	fmt.Printf("request %s\n", string(request[:n]))

	response := "hello " + string(request[:n])
	_, err = conn.Write([]byte(response))
	checkError(err)

}
