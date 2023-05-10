package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path, info.Name(), info.IsDir())
		return nil
	})

}
