package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())
}

// 生成n位随机数
func RandString(n int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}

func main() {
	//strings.Builder{}
	//strings.Reader{}
	// 定义strings.builder 结构体对象，类似在内存中的流对象(写文件对象)
	var builder strings.Builder
	builder.Write([]byte("我是kk\n")) // builder中写入字符切片
	builder.WriteString("我是阿宁\n")   // builder中写入字符串
	builder.WriteRune('b')          // builder中写入码点
	builder.WriteByte('b')          //  builder 中写入字节
	fmt.Println(builder.String())

	fmt.Println(builder.Len()) // 获取字节数量
	builder.Reset()            // 清空builder对象
	fmt.Println(builder.String())
	fmt.Println(builder.Len())

	fmt.Println(RandString(5))
	fmt.Println(RandString(6))
}
