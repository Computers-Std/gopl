package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launching Rocket...")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a sinle byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// Do nothing
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}

	launch()
}
