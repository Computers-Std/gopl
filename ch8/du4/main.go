package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
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

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			// Drain fileSizes before return
			for range fileSizes {
				// Do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes) // population progress
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Go(func() {
				walkDir(subdir, n, fileSizes)
			})
		} else {
			info, _ := entry.Info()
			fileSizes <- info.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.DirEntry {
	select {
	case sema <- struct{}{}: // access a token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
