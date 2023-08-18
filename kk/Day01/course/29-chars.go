package main

import "fmt"

func main() {
	letters := "abcdefghijklmn"

	for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}

	letters = "我爱中华人民共和国" // 重新赋值用 =
	for _, v := range letters {
		fmt.Printf("%q\n", v)
	}

}
