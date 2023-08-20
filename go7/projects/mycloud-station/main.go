package main

import (
	"fmt"
	"mycloud-station/cli"
	"os"
)

func main() {
	if err := cli.RootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
