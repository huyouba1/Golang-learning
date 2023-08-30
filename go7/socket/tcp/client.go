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
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err)

	request := "zcy"
	_, err = conn.Write([]byte(request))
	checkError(err)

	response := make([]byte, 256)
	n, err := conn.Read(response)
	checkError(err)
	fmt.Printf("client response: %s\n", string(response[:n]))

}
