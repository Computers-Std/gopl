// v2: Removed periodic output and using instant output, similar to du
// and Verbose by default.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type rootInfo struct {
	root int
	size int64
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel
	fileSizes := make(chan rootInfo)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Go(func() {
			walkDir(root, &n, i, fileSizes)
		})
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// var nfiles, nbytes int64
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

	for dir := range fileSizes {
		nfiles[dir.root]++
		nbytes[dir.root] += dir.size
	}
	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
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
