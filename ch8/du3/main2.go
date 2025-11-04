// The du1 command computes the disk usage of the files in a directory
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// walkDir recursively walks the file tree rooted at dir and sends the
// size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Go(func() {
				walkDir(subdir, n, fileSizes)
			})
		} else {
			info, _ := entry.Info() // NOTE ignoring errors
			fileSizes <- info.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	verbose := flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Go(func() {
			walkDir(root, &n, fileSizes)
		})
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}
