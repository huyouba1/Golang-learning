package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// 返回文件后缀
func fileExt(path string) string {
	pos := strings.LastIndex(path, ".")
	if pos < 0 {
		return ""
	}
	return path[strings.LastIndex(path, "."):]
}

func main() {
	path, _ := filepath.Abs("./1-filepath.go") // 获取绝对路径
	fmt.Println(path)
	fmt.Println(filepath.Base(path)) // 获取最后名称
	fmt.Println(filepath.Dir(path))  // 获取父目录

	fmt.Println(filepath.Clean("./../..////abc")) // 整理目录

	fmt.Println(filepath.Ext(path)) // 返回文件后缀
	fmt.Println(fileExt(path))

	fmt.Println(filepath.FromSlash("../../..../../asd/a///a")) // 拼路径
	fmt.Println(filepath.ToSlash("../../..../../asd/a///a"))

	//path2 := filepath.Dir(path)
	path2, _ := filepath.Abs("./test")
	fmt.Println(filepath.HasPrefix(path, path2))

	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs("."))

	dir, _ := filepath.Abs("/opt/cmdb")
	fmt.Println(filepath.Join(dir, "etc", "my.ini")) // 拼接路径

	fmt.Println("8888888888")
	fmt.Println(filepath.Split(path)) // 同时获取了dir和base

	paths := "/test/a:/test/b:/test/c" // 分割路径 ：linux  ; windows
	fmt.Println(filepath.SplitList(paths))

	fmt.Println(filepath.Glob("./test/*.go"))
	fmt.Println(filepath.Glob("./test/a.*")) // 按照要求格式进行匹配

	fmt.Println(filepath.Match("./test/a.go", "./test/a.go"))
	fmt.Println(filepath.Match("./test/a.*", "./test/a.go"))
	fmt.Println(filepath.Match("./test/b.*", "./test/a.go"))

	fmt.Println(strings.HasSuffix("asdasd.go", ".go"))

}
