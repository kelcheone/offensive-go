package main

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintf(w, "Executing /a")
	case "/b":
		fmt.Fprintf(w, "Executing /b")
	case "/c":
		fmt.Fprintf(w, "Executing /c")
	case "/d":
		fmt.Fprintf(w, "Executing /d")
	default:
		http.Error(w, "404 not found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8076", &r)
}
