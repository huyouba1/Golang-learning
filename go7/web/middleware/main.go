package main

import (
	"log"
	"net/http"
	"time"
)

var limitCh = make(chan struct{}, 100)

type middleware func(handler http.Handler) http.Handler

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		begin := time.Now()
		next.ServeHTTP(writer, request)
		timeElapsed := time.Since(begin)
		log.Printf("request %s use %d ms\n", request.URL.Path, timeElapsed.Milliseconds())
	})
}

func limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		limitCh <- struct{}{}
		log.Printf("current %d\n", len(limitCh))
		next.ServeHTTP(writer, request)
		<-limitCh
	})
}

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middlewareChain: make([]middleware, 0, 10),
		mux:             make(map[string]http.Handler, 10),
	}
}

// 添加中间件
func (self *Router) Use(m middleware) {
	self.middlewareChain = append(self.middlewareChain, m)
}

// 自定义路由
func (self *Router) Add(path string, handler http.Handler) {
	var mergeHandler = handler
	for i := 0; i < len(self.middlewareChain); i++ {
		mergeHandler = self.middlewareChain[i](mergeHandler)
	}
	self.mux[path] = mergeHandler
}

func get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	w.Write([]byte("How are you？"))
}

func main() {

	router := NewRouter()
	router.Use(timeMiddleware)
	router.Use(limitMiddleware)

	router.Add("/", http.HandlerFunc(get))

	for path, handler := range router.mux {
		http.Handle(path, handler)
	}

	//http.HandlerFunc("/", get)
	http.ListenAndServe(":5656", nil)
}
