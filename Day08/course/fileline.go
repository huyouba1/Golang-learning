package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 计算文件内容行数
func fileLine(path string) int {
	cnt := 0
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()

	// 带缓冲IO读取文件，按行读取，计算行数
	reader := bufio.NewReader(file)
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		// 过滤空行数据及注释
		txt := strings.TrimSpace(string(ctx))
		if txt == "" || strings.HasPrefix(txt, "//") {
			continue
		}

		cnt++
	}
	return cnt
}

func main() {
	dir := "../"
	total := 0

	var wg sync.WaitGroup

	channel := make(chan int, 10)

	// 遍历文件夹，计算每个go文件的行数并计算总数
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && ".go" == filepath.Ext(path) {
			wg.Add(1)
			go func() {
				/*
					问题：
					1. 主例程结束之后，工作例程还在执行
				*/
				cnt := fileLine(path)
				total += cnt
				channel <- cnt
				wg.Done()
			}()
		}
		return nil
	})
	//fmt.Println(fileLine("4-account.go"))
	wg.Wait()
	fmt.Println(total)

	// a. walk 之前 x
	// b. wait 之前
	// c. wait 之后
	// d. walk 之前,goroutine
	// b. wait 之前,goroutine
	// c. wait 之后,goroutine
}
