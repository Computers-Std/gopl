package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(r io.Reader, input chan<- string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		input <- s.Text()
	}
	close(input)
	if err := s.Err(); err != nil {
		log.Println("scan:", err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := make(chan string)
	go scan(c, input)

	wg := &sync.WaitGroup{}
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)

	for {
		select {
		case in, ok := <-input:
			// client closed connection
			if !ok {
				fmt.Println("client closed connection")
				wg.Wait() // wait for goroutines to finish before return/exit
				return
			}
			// got input, reset timeout
			if !timer.Stop() { // stop the timer, if already started or already fired
				<-timer.C // then, drain the channel if it had already fired
			}
			timer.Reset(timeout) // reset the timer anyways, to start a fresh countdown

			wg.Go(func() {
				echo(c, in, 1*time.Second)
			})
		case <-timer.C:
			fmt.Println(c, "Server timeout, closing connection.")
			wg.Wait()
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Echo server listening on localhost:8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn) // go ...
	}
}
