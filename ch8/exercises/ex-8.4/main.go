package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// modified reverb2
func handleConn(c *net.TCPConn) int {
	count := make(chan int)
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	defer c.Close() // close the conn when function returns

	// Create a goroutine for each line read from the connection
	for input.Scan() {
		wg.Add(1)
		go func(text string) { // Capture text argument to avoid closure issues
			defer wg.Done()
			echo(c, text, 1*time.Second)
			count <- 1
		}(input.Text())
	}

	// Handle the closing of the connection and channel
	go func() {
		wg.Wait()      // Wait for all goroutines to finish
		close(count)   // Close the channel safely after all goroutines are done
		c.CloseWrite() // Close the connection
	}()

	var total int
	for i := range count {
		total += i
	}
	return total
}

func main() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
