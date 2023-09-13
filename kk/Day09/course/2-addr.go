package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取网络地址
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		//fmt.Println(addr.Network(), addr.String())
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() && v.IP.To4() != nil {
				fmt.Println("本地IPv4地址：", v.IP.String())
			}
		case *net.IPAddr:
			if !v.IP.IsLoopback() && v.IP.To4() != nil {
				fmt.Println("本地IPv4地址：", v.IP.String())
			}
			//fmt.Println(addr.String())
		}
	}
}
