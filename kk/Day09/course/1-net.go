package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "6666"))
	fmt.Println(net.JoinHostPort("::1", "123"))
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	fmt.Println(net.SplitHostPort("[::1]:9999"))

	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupHost("www.baidu.com"))

	// IP结构体
	ip := net.ParseIP("110.242.68.3")
	fmt.Printf("%T, %#v\n", ip, ip)
	ip = net.ParseIP("::1")
	fmt.Printf("%T, %#v\n", ip, ip)

	hosts, err := net.LookupIP("www.baidu.com")
	fmt.Println(hosts, err)

	// IPNet IP范围，cidr格式
	ip, ipnet, err := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ip, ipnet, err)

	fmt.Println(ipnet.Contains(ip))
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.1")))
	//fmt.Println(net.ParseIP("192.168.2.x"))
	fmt.Println(ipnet.Network())
}
