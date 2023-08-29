package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Get() {
	resp, err := http.Get("http://127.0.0.1:5656")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}

func Post() {
	reader := strings.NewReader("hello server")
	resp, err := http.Post("http://127.0.0.1:5656/user/zxy/vip/bj/haidian", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
	for k, v := range resp.Header {
		fmt.Printf("%s = %v\n", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
}

func complexHttpRequest() {
	reader := strings.NewReader("hello server")
	req, err := http.NewRequest("POST", "http://127.0.0.1:8000", reader)
	if err != nil {
		panic(req)
	} else {
		req.Header.Add("User-Agent", "爱情")
		req.Header.Add("My-Header", "My-Header-Value")
		req.AddCookie(&http.Cookie{
			Name:  "auth",
			Value: "aaaa",
		})
		client := &http.Client{
			Timeout: 100 * time.Microsecond,
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.Body)
		}
	}
}

func main() {
	for i := 0; i < 200; i++ {
		go func() {
			Get()
		}()
	}

	time.Sleep(60 * time.Minute)
	//Get()
	//Post()
	//complexHttpRequest()
}
