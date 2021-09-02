package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func handleRequest(p string) {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":"+p, nil)
}

func main() {
	port := "5000"
	fmt.Println("Server listion on: http://localhost:" + port)
	handleRequest(port)
}
