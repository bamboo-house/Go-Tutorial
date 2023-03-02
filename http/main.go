package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Httpサーバー")
}

func main() {
	fmt.Println("hello warld")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}