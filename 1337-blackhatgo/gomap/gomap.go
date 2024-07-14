package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

func worker(ports, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			conn.Close()
			results <- p
		}
	}
}

func main() {
	const maxPort = 1024
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int
	var wg sync.WaitGroup

	// Increase the number of workers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(ports, results, &wg)
	}

	go func() {
		for i := 1; i <= maxPort; i++ {
			ports <- i
		}
		close(ports)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		openports = append(openports, port)
	}

	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
