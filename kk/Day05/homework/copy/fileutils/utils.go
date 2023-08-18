package fileutils

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) string {
	// 打开给定路径的文件
	file, err := os.Open(path)
	// 如果发生错误，则抛出异常
	if err != nil {
		panic(err)
	}
	// 延迟关闭文件
	defer file.Close()
	// 创建一个字节切片来存储文件内容
	txt := make([]byte, 0, 100)
	// 创建一个字节切片来读取文件内容
	cxt := make([]byte, 10)
	// 读取文件内容，直到文件结束
	for {
		n, err := file.Read(cxt)
		if err == io.EOF {
			break
		}
		// 将读取的字节追加到字节切片中
		txt = append(txt, cxt[:n]...)
	}
	// 将文件内容作为字符串返回
	return string(txt)
}

// 写文件函数，将给定的文本写入给定路径的文件
func WriteFile(path, txt string) {
	// 打开给定路径的文件
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	// 如果发生错误，则抛出异常
	if err != nil {
		panic(err)
	}
	// 延迟关闭文件
	defer file.Close()
	// 将给定的文本写入文件
	file.WriteString(txt)
}

// IsDir函数，检查给定路径是否为目录
func IsDir(path string) bool {
	// 获取给定路径的文件信息
	fileinfo, err := os.Stat(path)
	// 如果发生错误，则抛出异常
	if err != nil {
		panic(err)
	}
	// 返回文件信息是否为目录
	return fileinfo.IsDir()
}

// TraverseCopy函数，将源目录的内容复制到目标目录
func TraverseCopy(src, dest string) {
	// 打开源目录
	file, err := os.Open(src)
	// 如果发生错误，则返回
	if err != nil {
		return
	}
	// 延迟关闭源目录
	defer file.Close()
	// 读取源目录的内容
	fileInfos, err := file.Readdir(-1)
	// 如果发生错误，则返回
	if err != nil {
		return
	}
	// 遍历源目录的内容
	for _, fileInfo := range fileInfos {
		// 创建源路径和目标路径
		spath := src + "/" + fileInfo.Name()
		dpath := dest + "/" + fileInfo.Name()
		// 打印源路径
		fmt.Println(spath)
		// 打印目标路径
		fmt.Println(dpath)
		// 如果文件信息是目录
		if fileInfo.IsDir() {
			// 打印目录路径
			fmt.Println("Dir: ", spath)
			// 创建目标目录
			os.MkdirAll(dpath, os.ModePerm)
			// 递归复制源目录的内容到目标目录
			TraverseCopy(spath, dpath)
		} else {
			// 打印文件路径
			fmt.Println("File:", spath)
			// 从源路径复制文件到目标路径
			WriteFile(dpath, ReadFile(spath))
		}
	}

	fmt.Printf("%#v\n", fileInfos)
}
