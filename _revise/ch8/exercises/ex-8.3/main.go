// Exercise 8.3: In netcat3, the interface value conn has the concrete
// type *net.TCPConn, which represents a TCP connection. A TCP
// connection consists of two halves that may be closed independently
// using its CloseRead and CloseWrite methods. Modify the main
// goroutine of netcat3 to close only the write half of the connection
// so that the program will continue to print the final echoes from
// the reverb1 server even after the standard input has been closed.
// (Doing this for the reverb2 server is harder ; see Exercise 8.4.)
//// FROM: ukiran/gopl/ch8/exercises/ex-8.3/main.go

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		// Read from the server's connection
		io.Copy(os.Stdout, conn) // NOTE ignoring errors
		log.Println("Done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		if err == io.EOF {
			return
		}
		log.Fatal(err)
	}
}
