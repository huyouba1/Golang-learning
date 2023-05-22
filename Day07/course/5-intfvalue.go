package main

import "fmt"

// 定义有Send和senDall方法的接口
type Sender interface {
	Send(string, string) error
	SendAll([]string, string) error
}

// 定义只有Send方法的接口
type SingleSender interface {
	Send(string, string) error
}

// 定义结构体，并定义Send和SendAll方法
type EmailSender struct {
}

func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给：%s,内容：%s\n", to, msg)
	return nil
}

func (s EmailSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件给：%s,内容：%s\n", to, msg)
	return nil
}

func main() {
	// 定义sender接口和singleSender接口
	var sender Sender
	var singleSender SingleSender

	// 定义结构体对象
	emailSender := EmailSender{}

	// 将 emailSender赋值给sender对象
	sender = emailSender

	// 将Sender接口的对象赋值给singleSender接口对象
	singleSender = sender
	fmt.Printf("%T,%#v\n", sender, sender)
	fmt.Printf("%T,%#v\n", singleSender, singleSender)

	/*
		singleSender = emailSender
		// 不能赋值(没有SendAll方法)
		sender = singleSender
		fmt.Printf("%T,%#v\n", sender, sender)
		fmt.Printf("%T,%#v\n", singleSender, singleSender)
	*/
}
