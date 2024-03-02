package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in>")

	return os.Stdin.Read(b)
}

type FooWriter struct{}

/*
	func (fooWriter *FooWriter) Writer(b []byte) (int, error) {
		fmt.Print("out>")
		return os.Stdout.Write(b)
	}
*/
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
