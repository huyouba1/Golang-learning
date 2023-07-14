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
				//total += cnt
				channel <- cnt
				wg.Done()
			}()
		}
		return nil
	})

	//var wgTotal sync.WaitGroup
	//wgTotal.Add(1)
	exit := make(chan struct{})
	go func() {
		for cnt := range channel {
			total += cnt
		}
		//wgTotal.Done()
		exit <- struct{}{}
	}()

	//fmt.Println(fileLine("4-account.go"))
	//for <- channel
	wg.Wait()
	close(channel)

	//wgTotal.Wait()
	<-exit
	fmt.Println(total)

	// a. walk 之前 x
	// b. wait 之前 x
	// c. wait 之后 x
	// d. walk 之前,goroutine v
	// b. wait 之前,goroutine v
	// c. wait 之后,goroutine
}
