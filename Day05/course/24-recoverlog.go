package main

import (
	"log"
	"os"
)

func main() {
	logfile, err := os.OpenFile("appendlog1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(logfile)
	log.Println("test")
}
