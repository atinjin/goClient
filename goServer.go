package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Golang Server:%s", r.URL.Path[1:])
	fmt.Fprintf(w, "Your auth header is %s", r.Header.Get("Authorization"))
}

func main() {
	fmt.Print("Hello, Golang");
	http.HandleFunc("/test", handler)
	http.HandleFunc("/session", handler)

	http.ListenAndServe(":8080", nil)
}
