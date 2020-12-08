package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintf(w,h)
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8080"
	}
	http.HandlerFunc("/headers",headers)
	server.ListenAndServe()
}