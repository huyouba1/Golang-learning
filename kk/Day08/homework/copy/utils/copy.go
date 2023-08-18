package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// copy 文件
func CopyFile(src, dst string, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	dstFile.Close()

	srcBuf := bufio.NewScanner(srcFile)
	dstBuf := bufio.NewWriter(dstFile)
	dstBuf.Flush()

	for srcBuf.Scan() {
		content := srcBuf.Bytes()
		content = append(content, '\n')
		if _, err := dstBuf.Write(content); err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}

// Copy 目录
func CopyDir(src, dst string, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	doneLocal := make(chan struct{})
	err := os.Mkdir(dst, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	files, err := srcFile.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			go CopyFile(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()), doneLocal)
		} else {
			go CopyDir(filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()), doneLocal)
		}
	}

	for range files {
		<-doneLocal
	}
	close(doneLocal)
}

// copy 文件或目录
func Copy(src, dst string) {
	fmt.Println(time.Now())
	done := make(chan struct{})

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	if fileInfo.IsDir() {
		go CopyDir(src, dst, done)
	} else {
		go CopyFile(src, dst, done)
	}

	<-done
	fmt.Println(time.Now())
	close(done)
}
