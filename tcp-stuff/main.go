package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
	ports := make(chan int, 100)
	results := make(chan int)

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results

		if port != 0 {
			openPorts = append(openPorts, port)
		}

	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			return
		}
		conn.Close()
		results <- p

	}
}
