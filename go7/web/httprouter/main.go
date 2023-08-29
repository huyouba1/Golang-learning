package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func handle(method string, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request method %s\n", r.Method)
	fmt.Println("request body")
	io.Copy(os.Stdout, r.Body)

	//fmt.Fprintf(w, "hello boy")

	w.Write([]byte("hello boy your request method is " + r.Method))
}

func get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var arr []int
	_ = arr[1]
	handle("GET", w, r)
}

func post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	handle("POST", w, r)
}

func p1(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//var a = 0
	//_ = 4 / a
	var arr []int
	_ = arr[1]
}

func main() {
	router := httprouter.New()
	router.GET("/", get)
	router.POST("/", post)

	// restful
	router.POST("/user/:name/:type/*addr", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Printf("name=%s  type=%s   addr=%s \n", params.ByName("name"), params.ByName("type"), params.ByName("addr"))
	})

	// 返回静态文件  cli ->  localhost:5656/file/a.html
	router.ServeFiles("/file/*filepath", http.Dir("./static"))

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprintf(writer, "server panic %v", i)
	}
	//router.GET("/panic", p1)

	http.ListenAndServe(":5656", router)

}
