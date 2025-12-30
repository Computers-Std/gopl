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
	userlist := make(map[string]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				select {
				case cli.writer <- msg:
					// message sent successfully
				default:
					// means this channel is full/blocked. skip this
					// client!
				}
			}
		case request := <-entering:
			if userlist[request.cli.user] {
				request.replyCh <- false // name is taken
			} else {
				userlist[request.cli.user] = true
				clients[request.cli] = true
				request.replyCh <- true

				var curr strings.Builder
				curr.WriteString("currently online:")
				switch {
				case len(clients) > 1:
					for c := range clients {
						if c != request.cli {
							fmt.Fprintf(&curr, " %cs", c.user)
						}
					}
					request.cli.writer <- curr.String()
				default:
					curr.WriteString(" just you")
					request.cli.writer <- curr.String()
				}
			}
		case cli := <-leaving:
			delete(clients, cli)
			delete(userlist, cli.user)
			close(cli.writer)
		}
	}
}

type clientSession struct {
	conn    net.Conn
	input   *bufio.Scanner
	cli     client
	timeout time.Duration
}

func (cs *clientSession) handle() {
	messages <- cs.cli.user + " has arrived"
	cs.cli.writer <- "Welcome, " + cs.cli.user

	lines := make(chan string)
	go func() {
		for cs.input.Scan() {
			lines <- cs.input.Text()
		}
		close(lines)
	}()

	timer := time.NewTimer(cs.timeout)
	defer timer.Stop()

	for {
		select {
		case line, ok := <-lines:
			if !ok {
				leaving <- cs.cli
				messages <- cs.cli.user + " has left"
				return
			}
			if !timer.Stop() {
				select { // to drain
				case <-timer.C:
				default:
				}
			}
			timer.Reset(cs.timeout)
			messages <- cs.cli.user + ": " + line
		case <-timer.C:
			fmt.Fprintln(cs.conn, "You timed out!")
			cs.conn.Close() // The "Kick" to unblock input.Scan()
			leaving <- cs.cli
			messages <- cs.cli.user + " has left (timeout)"
			return
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

		// Give the user a buffer of 20 messages
		ch := make(chan string, 20)

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
			cs := &clientSession{
				conn:    conn,
				input:   input,
				cli:     cli,
				timeout: 25 * time.Second,
			}
			cs.handle()
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
