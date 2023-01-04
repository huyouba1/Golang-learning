package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l", "-t") //定义
	//bytes, err := cmd.Output()  // 执行
	//fmt.Println(string(bytes), err)
	output, _ := cmd.StdoutPipe()
	cmd.Start()
	io.Copy(os.Stdout, output)

	cmd.Wait()
}
