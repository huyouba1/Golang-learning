package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8888/?a=1&b=2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Proto, resp.StatusCode, resp.Status)
		fmt.Println(resp.Header)
		io.Copy(os.Stdout, resp.Body)
	}

	buffer := bytes.NewBufferString(`{"a":1}`)
	resp, err := http.Post("http://localhost:8888", "application/json", buffer)
	fmt.Println(resp, err)

	params := url.Values{}
	params.Add("a", "1")
	params.Add("a", "2")
	params.Add("b", "3")
	resp, err := http.PostForm("http://localhost:8888", params)
	fmt.Println(resp, err)
}
