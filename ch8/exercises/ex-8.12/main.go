// Exercise 8.12: Make the broadcaster announce the current set of
// clients to each new arrival. This requires that the clients set and
// the entering and leaving channels record the client name too.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			curr.WriteString("Currently online:")
			if len(clients) > 1 {
				for c := range clients {
					if c != cli {
						fmt.Fprintf(&curr, " %s", c.name)
					}
				}
			} else {
				curr.WriteString(" nil")
			}
			cli.out <- curr.String()

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

func handleConn(conn net.Conn) {
	who := conn.RemoteAddr().String()
	var cli client
	ch := make(chan string)

	go clientWriter(conn, ch)

	cli.out = ch
	cli.name = who

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE ignoring network errors
	}
}
