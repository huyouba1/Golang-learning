package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	//stdout := os.Stdout

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	io.Copy(os.Stdout, stdout)
	cmd.Wait()
}
