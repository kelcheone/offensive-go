package main

import (
	"io"
	"log"
	"net"
)

func handler(src net.Conn) {
	dst, err := net.Dial("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
		}
	}()
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
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
