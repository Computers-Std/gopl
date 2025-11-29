package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type Clock struct {
	loc, host string
}

var clocks []*Clock

func init() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall ZONE=HOST ...")
		os.Exit(1)
	}
	// Parse arguments and fill clocks slice
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "invalid arg: %s\n", arg)
			os.Exit(1)
		}
		clocks = append(clocks, &Clock{fields[0], fields[1]})
	}
}

func (c *Clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.loc, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Printf("can't read from %s: %s", c.loc, err)
	}
	fmt.Println(c.loc, "done")
}

func main() {
	// rest calls : Non-Blocking
	if len(clocks) > 1 {
		for _, c := range clocks[1:] {
			conn, err := net.Dial("tcp", c.host)
			if err != nil {
				log.Println("Error connection to", c.loc, err)
			}
			defer conn.Close()
			go c.watch(os.Stdout, conn)
		}
	}

	// first call : Blocking
	conn, err := net.Dial("tcp", clocks[0].host)
	if err != nil {
		log.Println("Error connection to", clocks[0].loc, err)
		return
	}
	defer conn.Close()
	clocks[0].watch(os.Stdout, conn)
}
