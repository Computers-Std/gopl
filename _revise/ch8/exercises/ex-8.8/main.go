package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	lines := make(chan string, 1) // Buffer prevents sender hang

	// concurrent read
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			lines <- input.Text()
		}
		// this will be triggered, once the client is closed
		close(lines)
		fmt.Println("Scanner exiting")
	}()

	for {
		timer := time.NewTimer(10 * time.Second)
		defer timer.Stop()

		select {
		case line := <-lines:
			timer.Stop()
			echo(c, line, 1*time.Second)
		case <-timer.C:
			fmt.Fprintln(c, "Closing connection.")
			return
		}
	}
}

/**
 * NOTE:
 * 1. A channel can only be closed by its sender, and not the receiver.
 * 2. Unbuffered channels can cause Goroutine Leaks if the receiver quits.
 */
