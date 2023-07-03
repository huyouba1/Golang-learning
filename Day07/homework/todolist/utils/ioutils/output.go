package ioutils

import "fmt"

func Error(prompt string) {
	fmt.Printf("[x] %s\n", prompt)
}
