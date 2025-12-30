// Exercise 8.14: Change the chat server’s network protocol so that
// each client provides its name on entering. Use that name instead of
// the network address when prefixing each message with its sender’s
// identity.

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
	user   string
	addr   string
	writer chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

var (
	users = make(map[string]bool)
	mu    sync.Mutex
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.writer <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			var curr strings.Builder
			curr.WriteString("currently online:")
			if len(clients) > 1 {
				for c := range clients {
					if c != cli {
						fmt.Fprintf(&curr, " %s", c.user)
					}
				}
				cli.writer <- curr.String()
			}
		case cli := <-leaving:
			mu.Lock()
			delete(users, cli.user)
			mu.Unlock()
			delete(clients, cli)
			close(cli.writer)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	username := getUserName(conn)
	if username == "" {
		return
	}
	addr := conn.RemoteAddr().String()
	ch := make(chan string)
	cli := client{user: username, addr: addr, writer: ch}

	go clientWriter(conn, ch)

	cli.writer <- "You are " + username
	messages <- username + " has arrived"
	entering <- cli

	lines := make(chan string, 1)
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			lines <- input.Text()
		}
		close(lines)
	}()

	idleTimeout := 20 * time.Second
	timer := time.NewTimer(idleTimeout)
	defer timer.Stop()

	for {
		select {
		case line, ok := <-lines:
			if !ok {
				leaving <- cli
				messages <- cli.user + " has left"
				return
			}
			timer.Reset(idleTimeout)
			messages <- cli.user + ": " + line
		case <-timer.C:
			cli.writer <- "closing connection..."
			leaving <- cli
			messages <- cli.user + " has left"
			conn.Close() // this will cause input.Scan() to fail
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func getUserName(conn net.Conn) string {
	input := bufio.NewScanner(conn)
	for {
		fmt.Fprint(conn, "Enter you name: ")
		if !input.Scan() {
			return ""
		}
		username := strings.TrimSpace(input.Text())
		if username == "" {
			fmt.Fprintln(conn, "Please enter a valid name.")
			continue
		}
		mu.Lock()
		_, exists := users[username]
		if exists {
			mu.Unlock()
			fmt.Fprintf(conn, "Error: Name '%s' is already taken!\n", username)
			continue
		}
		users[username] = true
		mu.Unlock()
		return username
	}
}
