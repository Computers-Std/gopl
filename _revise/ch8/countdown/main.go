// countdown1, countdown2
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launching Rocket...OO_OO_OO")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a sinle byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	select {
	case <-time.After(10 * time.Second):
	// Do nothing
	case <-abort:
		fmt.Println("Launch aborted")
		return
	}

	launch()
}
