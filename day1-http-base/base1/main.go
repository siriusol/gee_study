package main

import (
	"fmt"
	"log"
	"net/http"
)

// indexHandler echos req.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path=%s\n", req.URL.Path)
	fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path) // %q 会使值被 "" 包裹
}

// helloHandler echos req.URL.Header
// todo pr: fix r.URL.Header to r.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}

func main() {
	// 调换两个 handler 的顺序并不会改变匹配优先级
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
