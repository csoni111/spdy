package main

import (
	"fmt"
	"github.com/nileshjagnik/spdy"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	spdy.EnableDebug()
	err := spdy.ListenAndServe("localhost:4040", nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}
