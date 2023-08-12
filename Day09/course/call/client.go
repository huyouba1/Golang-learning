package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

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
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	func() {
		defer conn.Close()

		log.Printf("Connected")

		reader := bufio.NewReader(conn)
		for {
			ctx := Input("Client 请输出你要回复的内容:")
			fmt.Fprintf(conn, ctx+"\n")
			log.Println("等待服务端返回数据....")
			line, _, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("服务端响应： %s\n", string(line))
		}
	}()
}
