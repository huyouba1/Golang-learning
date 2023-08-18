package main

func main() {
	// 定义管道，元素为int类型
	var channel chan int
	channel = make(chan int)
	<-channel
}
