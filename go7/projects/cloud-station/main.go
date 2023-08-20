package main

import (
	"cloud-station/cli"
	"fmt"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
