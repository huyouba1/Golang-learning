package main

import "fmt"

type Sender interface {
	Send(to, msg string) error
}

// 定义EmailSender结构体，并实现Sender接口
type EmailSender struct {
}

func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送邮件")
	return nil
}

// 定义EmailSender结构体，并实现Sender接口
type SMSSender struct {
}

func (s SMSSender) Send(to, msg string) error {
	fmt.Println("发送短信")
	return nil
}

func main() {
	var sender Sender
	sender = SMSSender{}
	sender.Send("", "") // 调用SMSSender

	sender = EmailSender{}
	sender.Send("", "") // 调用EmailSender

	// 把不同的类型的变量赋值给了同一个变量
}
