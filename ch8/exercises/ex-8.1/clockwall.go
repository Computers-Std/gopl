package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func parse(arg string) ([]string, error) {
	arr := strings.Split(arg, "=")
	if len(arr) != 2 {
		return nil, errors.New("invalid input")
	}
	return arr, nil
}

func main() {
	for _, zone := range os.Args[2:] {
		arr, err := parse(zone)
		if err != nil {
			log.Fatal(err)
		}
		loc, addr := arr[0], arr[1]
		go dial(loc, addr)
	}

	// [24-10-2025] NOTE: Why to do this seperately instead of using
	// os.Args[1:] in above loop

	// main goroutine
	arr, err := parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dial(arr[0], arr[1])
}

func mustCopy(dst io.Writer, src io.Reader, location string) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func dial(loc, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("Error connecting to", loc, err)
		return
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn, loc)
}

// Usage: clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
