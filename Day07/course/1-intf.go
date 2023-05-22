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

// 值接受者
func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件给：%s,内容：%s", to, msg)
	return nil
}

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

	// 给接口变量sender赋值
	// 创建EmailSender对象
	emailSender := EmailSender{"smtp.qq.com", 456, "kk", "123"}

	sender = emailSender
	fmt.Printf("%T,%#v\n", sender, sender)

	// 接口对象能不能访问属性
	//fmt.Println(sender.addr)  接口对象是不能访问属性的

	// 接口只能调用声明过的方法
	sender.Send("哈诶exu", "做业绩高")
	sender.SendALL([]string{"哈诶exu", "asdasda"}, "做业绩高")
	//sender.SendCc("asdasd", "ddddd", "结果")   // 只能定义接口定义好的对象

	pemailSender := &EmailSender{"smtp.qq.com", 456, "kk", "123"}
	sender = pemailSender
	fmt.Printf("%T,%#v\n", sender, sender)
}
