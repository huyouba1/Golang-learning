package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"net/http"
)

var (
	proxyaddr string
	upstream  []string
)

func init() {
	flag.StringVar(&proxyaddr, "proxy", "127.0.0.1:8888", "proxy's ip:port")
}

func main() {
	flag.Parse()
	upstream = []string{"127.0.0.1:9999"}
	http.HandleFunc("/proxy/", func(resp http.ResponseWriter, req *http.Request) {
		reader := bufio.NewReader(req.Body)
		rinfo := make([]byte, 1024)
		n, _ := reader.Read(rinfo)
		for _, endpoint := range upstream {
			conn, err := net.Dial("tcp", endpoint)
			if err != nil {
				log.Println(err)
				continue
			}
			func(conn net.Conn) {
				defer conn.Close()
				conn.Write(rinfo[:n])
			}(conn)
		}
	})
	http.ListenAndServe(proxyaddr, nil)
}
