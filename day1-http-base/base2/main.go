package main

import (
	"fmt"
	"net/http"
)

// todo pr 代码的运行结果与之前的是一致的。 这句话并不准确，比如 404 的处理，内置的处理器是没有 404 的。

// Engine is the unique handler for all requests
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found, URL:%s\n", req.URL)
		fmt.Fprintf(w, "404 Not Found, URL:%q\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":8080", engine)
}
