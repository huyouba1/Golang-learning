package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 创建一个临时文件夹，父目录为./test，文件夹命名以log为前缀+随机数字
	// 返回文件夹目录路径和错误error
	dir, _ := ioutil.TempDir("./test", "log")
	file, _ := os.Create(filepath.Join(dir, "1.log"))

	file.WriteString(time.Now().Format("2006-01-02"))
	file.Close()
	os.RemoveAll(dir)
}
