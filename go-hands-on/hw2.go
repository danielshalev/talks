package main

import (
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", sayHello)	
	http.ListenAndServe(":8085", nil)
}