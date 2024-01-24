package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 10000)
	results := make(chan int)
	var openports []int
	for i := 1; i <= cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	for port := range openports {
		fmt.Printf("%d is open\n", port)
	}

}
