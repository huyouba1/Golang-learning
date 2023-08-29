package main

import (
	"fmt"
	"net/http"
)

func BoyHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "boy")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "wocao")
		for k, v := range request.Header {
			fmt.Printf("%s == %v\n", k, v)
		}
	})

	http.HandleFunc("/boy", BoyHandler)

	http.ListenAndServe("127.0.0.1:8000", nil)
}
