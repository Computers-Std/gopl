package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	user   string
	addr   string
	writer chan<- string
}

type registration struct {
	cli     client
	replyCh chan bool
}

var (
	entering = make(chan registration)
	leaving  = make(chan client)
	messages = make(chan string)
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

func broadcaster() {
	clients := make(map[client]bool)
	usernames := make(map[string]bool)

	for {
		select {
		case request := <-entering:
			if usernames[request.cli.user] {
				request.replyCh <- false // name is taken
			} else {
				usernames[request.cli.user] = true
				clients[request.cli] = true
				request.replyCh <- true

				var curr strings.Builder
				curr.WriteString("currently online:")
				switch {
				case len(clients) > 1:
					for c := range clients {
						if c != request.cli {
							fmt.Fprintf(&curr, " %s", c.user)
						}
					}
					request.cli.writer <- curr.String()
				default:
					curr.WriteString(" just you")
					request.cli.writer <- curr.String()
				}
			}
		case msg := <-messages:
			for cli := range clients {
				cli.writer <- msg
			}
		case cli := <-leaving:
			delete(clients, cli)
			delete(usernames, cli.user)
			close(cli.writer)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	var username string
	for {
		fmt.Fprintf(conn, "Enter name: ")
		if !input.Scan() {
			return
		}
		username = strings.TrimSpace(input.Text())
		if username == "" {
			continue
		}

		ch := make(chan string)
		cli := client{
			user:   username,
			addr:   conn.RemoteAddr().String(),
			writer: ch,
		}
		reply := make(chan bool)
		entering <- registration{cli: cli, replyCh: reply}

		if <-reply {
			// success!
			go clientWriter(conn, ch)
			handleSession(conn, cli, input)
			return
		}
		fmt.Fprintln(conn, "Name already taken.")
	}
}

// If you only return from handleSession on a timeout, the
// input.Scan() goroutine might stay alive for a short while longer
// until the defer in handleConn eventually closes the socket.
//
// By passing conn to handleSession, you can call conn.Close() the
// instant the timer expires. This forces the scanner goroutine to
// exit immediately.
func handleSession(conn net.Conn, cli client, input *bufio.Scanner) {
	messages <- cli.user + " has arrived"
	cli.writer <- "Welcome, " + cli.user

	lines := make(chan string)
	go func() {
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
			if !timer.Stop() {
				select { // to drain
				case <-timer.C:
				default:
				}
			}
			timer.Reset(idleTimeout)
			messages <- cli.user + ": " + line
		case <-timer.C:
			fmt.Fprintln(conn, "You timed out!") // Direct communication
			conn.Close()                         // The "Kick" to unblock input.Scan()
			leaving <- cli
			messages <- cli.user + " has left (timeout)"
			return // handleConn can also take care of conn.Close()
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
