package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type rootInfo struct {
	root int
	size int64
}

// walkDir recursively walks the file tree rooted at dir and sends the
// size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, root int, fileSizes chan<- rootInfo) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Go(func() {
				walkDir(subdir, n, root, fileSizes)
			})
		} else {
			info, _ := entry.Info() // NOTE ignoring errors
			fileSizes <- rootInfo{root, info.Size()}
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

func main() {
	verbose := flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan rootInfo)
	var n sync.WaitGroup
	for i, root := range roots {
		// n.Add(1)
		// go walkDir(root, &n, i, fileSizes)
		n.Go(func() {
			walkDir(root, &n, i, fileSizes)
		})
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles, nbytes := make([]int64, len(roots)), make([]int64, len(roots))

loop:
	for {
		select {
		case dir, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles[dir.root]++
			nbytes[dir.root] += dir.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n",
			nfiles[i], float64(nbytes[i])/1e9, r)
	}
}
