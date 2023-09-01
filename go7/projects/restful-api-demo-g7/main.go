package main

import (
	"fmt"
	"gitee.com/go-learn/restful-api-demo-g7/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
