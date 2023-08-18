package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "p", "", "path")
	flag.Parse()

	if name == "" {
		return
	}
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	hasher := md5.New()

	ctx := make([]byte, 1024)
	for {
		n, err := file.Read(ctx)
		if err == io.EOF {
			break
		}
		//ctx[:n]
		hasher.Write(ctx[:n])
	}
	fmt.Printf("%x", hasher.Sum(nil))
}

/*
$ go run 20-md5file.go -p 20-md5file.go
dfb6792602dd58e474c430fe0e2a5f84%

$ md5 20-md5file.go
MD5 (20-md5file.go) = dfb6792602dd58e474c430fe0e2a5f84

*/
