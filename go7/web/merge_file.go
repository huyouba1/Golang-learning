package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var fileChan = make(chan string, 10000)

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 构建FileReader
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					fileChan <- line
				}
				break
			} else {
				fmt.Println(err)
				break
			}
		} else {
			fileChan <- line
		}
	}
}

func WriteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	for {
		line := <-fileChan
		writer.WriteString(line)

	}
	writer.Flush()

}

func main() {

	for i := 1; i <= 3; i++ {
		fileName := "./" + strconv.Itoa(i)
		go readFile(fileName)

	}
}
