package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	listeAddr  string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listeAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for the websocket connection")
	flag.Parse()

	var err error

	jsTemplate, err = template.ParseFiles("logger.js")

	if err != nil {
		panic(err)
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}

	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", serveWS)
	mux.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
