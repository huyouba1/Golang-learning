package ioutils

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"os"
	"strings"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner((os.Stdin))
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func Password(prompt string) string {
	fmt.Print(prompt)
	if ctx, err := gopass.GetPasswd(); err != nil {
		return ""
	} else {
		return strings.TrimSpace(string(ctx))
	}
}
