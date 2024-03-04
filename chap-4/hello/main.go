package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s\n", r.URL.Query().Get("name"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello/", hello)
	mux.HandleFunc("GET /root/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "We are in the root!!")
	})
	fmt.Println("Listening on port :8076")
	http.ListenAndServe(":8076", mux)
}
