// Exercise 8.11: Following the approach of mirroredQuery in Section
// 8.4.4, implement a variant of fetch that requests several URLs
// concurrently. As soon as the first response arrives, cancel the
// other requests.

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
)

func fetch(ctx context.Context, url string,
	ch chan<- *http.Response, errs chan<- error,
) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		errs <- err
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errs <- err
		return
	}
	select {
	case ch <- resp:
	case <-ctx.Done():
		resp.Body.Close()
	}
}

func main() {
	flag.Parse()
	urls := flag.Args()
	if len(urls) == 0 {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	responses := make(chan *http.Response)
	errs := make(chan error)

	for _, url := range urls {
		go fetch(ctx, url, responses, errs)
	}

	var errCount int
	for {
		select {
		case firstResp := <-responses:
			cancel() // cancel others
			defer firstResp.Body.Close()
			fmt.Printf("Winner: %s\n", firstResp.Request.URL)
			return
		case <-errs:
			errCount++
			if errCount == len(urls) {
				fmt.Println("All requests failed.")
				return
			}
		}
	}
}
