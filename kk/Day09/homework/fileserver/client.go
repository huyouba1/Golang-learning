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

func ls(conn net.Conn) {
	fmt.Fprintf(conn, "ls|0|")
	reader := bufio.NewReader(conn)

	// 读取流中接收到的字符串，输出第一个"|"之前的内容
	sizeText, err := reader.ReadString('|')
	if err != nil {
		log.Printf("reader size err: %s ", err)
	}

	// 输出不包含"|"的内容的大小
	size, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Printf("strconv atoi has err: %s", err)
	}

	for size > 0 {
		name, err := reader.ReadString(':')
		if err != nil {
			log.Printf("size reader string err: %s", err)
		}
		fmt.Println(name[:len(name)-1])
		size--
	}
}

func cat(conn net.Conn, name string) {
	fmt.Fprintf(conn, "cat|1|%s|", name)

	reader := bufio.NewReader(conn)
	sizeText, err := reader.ReadString('|')
	if err != nil {
		log.Printf("cat sizeText err: %s", err)
	}

	size, err := strconv.Atoi(sizeText[:len(sizeText)-1])
	if err != nil {
		log.Printf("cat scrconv err: %s", err)
	}
	if size > 0 {
		ctx := make([]byte, size)
		n, err := reader.Read(ctx)
		log.Printf("cat read content  err %v", err)
		fmt.Printf("文件内容: %s", string(ctx[:n]))
	} else {
		fmt.Println("文件内容为空")
	}
}

func upload(conn net.Conn, filename string) {
	fmt.Fprintf(conn, "upload|1|%s|", filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("upload osopenfile err: %s", err)
	}
	defer file.Close()
	io.Copy(conn, file)
	log.Println("File sent:", filename)
}

func main() {
	logfile, err := os.OpenFile("client.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	log.SetOutput(logfile)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	addr := "localhost:8888"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Print("connected fileserver....")

	scanner := bufio.NewScanner(os.Stdin)
END:
	for {
		fmt.Print("请输入指令: ")
		scanner.Scan()
		input := scanner.Text()
		cmds := strings.Split(input, " ")
		switch cmds[0] {
		case "quit":
			fmt.Fprintf(conn, "quit|0|")
			break END
		case "cat":
			cat(conn, cmds[1])
		case "upload":
			upload(conn, cmds[1])
		case "ls":
			ls(conn)

		default:
			fmt.Println("输入指令错误")
		}
	}

}
