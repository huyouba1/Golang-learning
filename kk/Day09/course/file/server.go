package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func cmd(conn net.Conn) (string, []string) {
	reader := bufio.NewReader(conn)

	op, err := reader.ReadString('|')
	log.Printf("read op err: %s\n", err)
	// err 错误时退出
	if err != nil {
		return "quit", nil
	}

	cntText, err := reader.ReadString('|')
	log.Printf("read cntText err: %s\n", err)

	cnt, err := strconv.Atoi(cntText[:len(cntText)-1])
	log.Printf("atoi err %s\n", err)

	args := make([]string, 0, cnt)
	for cnt > 0 {
		param, err := reader.ReadString('|')
		log.Printf("read param err %s\n", err)
		args = append(args, param[:len(param)-1])
		cnt--
	}
	return op[:len(op)-1], args
}

func cat(conn net.Conn, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(conn, "0|")
	} else {
		defer file.Close()
		ctx := make([]byte, 1024)
		n, _ := file.Read(ctx)

		fmt.Fprintf(conn, "%d|%s", n, string(ctx[:n]))
	}

}

func ls(conn net.Conn) {
	file, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	names, err := file.Readdirnames(-1)

	// names 空  0|
	// names >0  1| name:

	suffix := ""
	if len(names) > 0 {
		suffix = ":"
	}
	fmt.Fprintf(conn, "%d|%s%s", len(names), strings.Join(names, ":"), suffix)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
END:
	for {
		op, args := cmd(conn)
		fmt.Println(op, args)
		switch op {
		case "ls":
			ls(conn)
		case "cat":
			cat(conn, args[0])
		case "quit":
			break END
		}
	}
	log.Printf("Client closed: %s\n", conn.RemoteAddr())
}

func main() {

	logfile, _ := os.OpenFile("fileserver.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	defer logfile.Close()

	// Todo: 带缓冲的io
	log.SetOutput(logfile)

	addr := "0.0.0.0:8888" // 监听所有网卡的9999端口  :9999

	// 启动监听服务
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	//延迟关闭
	defer listener.Close()
	log.Printf("listen on: %s", addr)

	for {
		// 接受客户端请求
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accept: %s", err)
			continue
		}

		log.Printf("client connected: %s", conn.RemoteAddr())
		// 处理客户端连接
		go handleConn(conn)
	}
}
