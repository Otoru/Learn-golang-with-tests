package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet sends a custom greeting to writer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

// GreeterHandler send a greet over HTTP
func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	const address = "127.0.0.1:5000"
	handler := http.HandlerFunc(GreeterHandler)
	err := http.ListenAndServe(address, handler)

	if err != nil {
		fmt.Println(err)
	}
}
