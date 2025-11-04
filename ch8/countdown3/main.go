package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// "abort" channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	// "nominal" channel
	tick := time.Tick(1 * time.Second)

	// multiplexing
	fmt.Println("Commencing countdown. Press return to abort.")
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
