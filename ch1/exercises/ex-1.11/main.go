// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		f, err := os.Create(strings.Join([]string{prettyURL(url, false), "html"}, "."))
		if err != nil {
			panic(err)
		}
		go fetch(prettyURL(url, true), ch, f) // start a goroutine
		defer f.Close()
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resource
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func prettyURL(url string, fullPath bool) string {
	prefix1 := "http://"
	prefix2 := "https://"

	if strings.HasPrefix(url, prefix1) || strings.HasPrefix(url, prefix2) {
		if fullPath == true {
			return url
		} else {
			return strings.SplitAfter(url, "://")[1]
		}
	} else {
		if fullPath == true {
			return strings.Join([]string{prefix1, url}, "")
		} else {
			return url
		}
	}
}
