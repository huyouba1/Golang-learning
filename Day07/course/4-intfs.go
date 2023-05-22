package main

import "fmt"

// 定义接口 Sender
type Sender interface {
	// 定义接口的行为(方法名称、参数列表、返回值列表)
	Send(to string, msg string) error
	SendALL([]string, string) error
}

type EmailSender struct {
	addr     string
	port     int
	user     string
	password string
}

// 指针接受者
func (s *EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给：%s,内容：%s", to, msg)
	return nil
}

// 值接受者
func (s EmailSender) SendALL(tos []string, msg string) error {
	fmt.Printf("发送邮件给：%#v,内容：%s\n", tos, msg)
	return nil
}

func (s EmailSender) SendCC(to string, cc string, msg string) error {
	fmt.Printf("发送邮件给：%s,抄送：%s,内容：%s", to, cc, msg)
	return nil
}

func main() {
	// 定义接口变量
	var sender Sender
	fmt.Printf("%T,%#v\n", sender, sender)

	// 定义指针对象 并进行赋值
	//var emailSender = EmailSender{}
	//sender = emailSender
	//fmt.Printf("%T,%#v\n", sender, sender)

	var pemailSender = &EmailSender{}
	sender = pemailSender
	fmt.Printf("%T,%#v\n", sender, sender)

}
