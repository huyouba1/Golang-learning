package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
1. flag h  help
2. 判断是否有参数，没有输出help 并推出
3. 参数即为path，设置变量
4. 判断文件是否存在，存在的话就打开，不存在就说没有找到，退出
4. 如果存在 bufio reader输出
*/

var h bool
var help bool

func traceFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, _ := reader.ReadString('\n')
		if strings.TrimSpace(line) != "" {
			fmt.Println(strings.TrimSpace(line))
		}
	}

}

// 判断文件是否存在
func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func main() {
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.BoolVar(&help, "help", false, "帮助信息")

	flag.Usage = func() {
		fmt.Println("tailf [file]  输出文件内容")
		flag.PrintDefaults()
	}
	flag.Parse()

	if h || help || flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	path := flag.Arg(0)
	if FileIsExists(path) {
		fmt.Printf("file: %s\n", path)
		traceFile(path)

	} else {
		fmt.Printf("file:[%s] not find!", path)
		os.Exit(1)
	}

}
