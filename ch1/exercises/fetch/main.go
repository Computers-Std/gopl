/*
   Exercise 1.7: Use io.Copy instead of io.ReadAll
   Exercise 1.8: Add prefix http:// if needed
   Exercise 1.9: Print HTTP status code in resp.Status
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		fmtUrl := urlChecked(url)
		resp, err := http.Get(fmtUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", fmtUrl, err)
			os.Exit(1)
		}
		status := resp.Status
		fmt.Printf("HTTP Status: %v\n", status)
	}
}

func urlChecked(url string) string {
	prefix := "http://"
	if strings.HasPrefix(url, prefix) {
		return url
	} else {
		return strings.Join([]string{prefix, url}, "")
	}
}
