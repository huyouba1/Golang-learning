package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// 定义一个用于返回命令和参数的函数
func cmd(conn net.Conn) (string, []string) {
	// 创建一个reader，接受conn的流
	reader := bufio.NewReader(conn)

	// 读取流中的字符串，返回第一个'|'之前的内容，也就是操作action|parm1|parm2|parm3
	op, err := reader.ReadString('|')
	if err != nil {
		log.Printf("read op err: %s\n", err)
		return "quit", nil
	}

	// 读取流中的字符串，返回第2个'|'之前的内容
	cntText, err := reader.ReadString('|')
	if err != nil {
		log.Printf("read cntText err: %s\n", err)
	}

	// 返回第2个|之前的长度，不包含|
	cnt, err := strconv.Atoi(cntText[:len(cntText)-1])
	if err != nil {
		log.Printf("atoi err %s\n", err)
	}

	// 创建一个和返回参数长度一样的切片
	args := make([]string, cnt)

	for cnt > 0 {
		parms, err := reader.ReadString('|')
		if err != nil {
			log.Printf("parms err %s\n", err)
		}
		args = append(args, parms[:len(parms)-1])
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

// 列出文件列表
func ls(conn net.Conn) {
	file, err := os.Open(".")
	if err != nil {
		log.Printf("ls osOpen has err: %s", err)
	}
	defer file.Close()

	// 将当前文件夹下面的文件名称全部读取出来，不会递归读取
	names, err := file.Readdirnames(-1)
	//fmt.Println(names)
	if err != nil {
		log.Printf("ls readdirname has err: %s", err)
	}

	// names 空  0|
	// names >0  1| name:

	suffix := ""
	// 设置后缀，如果有多个参数，则为:
	if len(names) > 0 {
		suffix = ":"
	}
	fmt.Fprintf(conn, "%d|%s%s", len(names), strings.Join(names, ":"), suffix)
	//fmt.Printf("%d|%s%s", len(names), strings.Join(names, ":"), suffix)
}

func upload(conn net.Conn, filename string) {
	// 接受文件名
	//_, err := conn.Read(filename)
	//if err != nil {
	//	log.Printf("upload err: %s", err)
	//}

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("upload create file err: %s ", err)
	}
	defer file.Close()

	io.Copy(file, conn)
	log.Println("File received and saved:", filename)
}

func handlecon(conn net.Conn) {
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
		case "upload":
			upload(conn, args[0])
		case "quit":
			//ls(conn)
			break END
		}
	}
	log.Printf("Client closed: %s\n", conn.RemoteAddr())
}

func main() {
	// 打开或创建一个日志文件
	logfile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	// 设置将日志内容标准输出到日志文件当中
	log.SetOutput(logfile)
	if err != nil {
		log.Fatal("os.open error: ", err)
	}
	defer logfile.Close()

	addr := ":8888"
	// 启动监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("net listen error: ", err)
	}
	defer listener.Close()

	log.Printf("listen on: %s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener accept error :", err)
			continue
		}

		log.Printf("client conneceted: %s", conn.RemoteAddr())
		handlecon(conn)
	}

}
