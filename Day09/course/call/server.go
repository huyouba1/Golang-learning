package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

//const dateFormat = "2006-01-02 15:04:05"

func Input(prompt string) string {
	/*
		使用带缓冲的io，可以使用下面任意的一种方式
		1. bufio.Scanner{}
		2. bufio.NewReader()
	*/

	// 打印调用函数时传入的参数(打印传入的内容)
	fmt.Print(prompt)
	// 创建一个读io缓冲
	scanner := bufio.NewScanner(os.Stdin)
	// 扫描输入的内容
	scanner.Scan()

	return scanner.Text()
}

func main() {
	addr := "127.0.0.1:9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Listen on [%s]\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		func() {
			defer conn.Close()
			log.Printf("Client [%s] is connected...", conn.RemoteAddr())

			reader := bufio.NewReader(conn)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					log.Fatal(err)
					break
				} else {
					if string(line) == "quit" {
						break
					}
					fmt.Printf("接受到数据：%s", string(line))
					ctx := Input("Server 请输出你要发送的内容:")
					fmt.Fprintf(conn, ctx+"\n")
					log.Println("等待客户端返回数据...")
				}
			}

		}()
	}

}
