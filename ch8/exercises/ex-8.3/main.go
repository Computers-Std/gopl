package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Define the server address to connect to
	addr := &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"), // Localhost
		Port: 8000,
	}

	// Establish a TCP connection to the address
	conn, err := net.DialTCP("tcp", nil, addr) // nil for localAddr means no specific source address
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer conn.Close() // Ensure connection is closed when main finishes

	// Create a channel to wait for the go routine to finish
	done := make(chan struct{})

	// Launch a goroutine to copy data from the connection to stdout
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Println("Error copying data:", err)
		}
		log.Println("Done reading from connection")
		done <- struct{}{} // Signal when done
	}()

	// Copy from stdin to the connection
	mustCopy(conn, os.Stdin)

	// Gracefully close the write side of the connection
	err = conn.CloseWrite()
	if err != nil {
		log.Fatal("Error closing write side of the connection:", err)
	}

	// Wait for the goroutine to finish
	<-done
}
