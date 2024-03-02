package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data")
			break
		}

		log.Printf("Received %d bytes: %s\n", len(s), s)

		log.Println("Writing data")

		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		writer.Flush()
	}
}

func echo2(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Could not read or write data")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	log.Println("Listening on port: 20080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connections")
		}

		log.Println("Received connection")

		// go echo(conn)
		go echo2(conn)
	}
}
