// Exercise 8.13: Make the chat server disconnect idle clients, such
// as those that have sent no messages in the last five minutes. Hint:
// calling conn.Close() in another goroutine unblocks active Read
// calls such as the one done by input.Scan().
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	out  chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.out <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var curr strings.Builder
			curr.WriteString("currently online:")
			if len(clients) > 1 {
				for c := range clients {
					if c != cli {
						fmt.Fprintf(&curr, " %s", c.name)
					}
				}
				cli.out <- curr.String()
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	who := conn.RemoteAddr().String()
	ch := make(chan string)
	cli := client{name: who, out: ch}

	go clientWriter(conn, ch)

	cli.out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	lines := make(chan string, 1)
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			lines <- input.Text()
		}
		close(lines) // triggered, once the client is closed
	}()

	idleTimeout := 10 * time.Second
	timer := time.NewTimer(idleTimeout)
	defer timer.Stop()

	for {
		select {
		case line, ok := <-lines:
			if !ok {
				leaving <- cli
				messages <- who + " has left"
				return
			}
			timer.Reset(idleTimeout)
			messages <- who + ": " + line

		case <-timer.C:
			cli.out <- "closing connection..."
			leaving <- cli
			messages <- who + " has left"

			// By given Hint: unblocks active Read calls such as the one
			// done by input.Scan().
			conn.Close()
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
