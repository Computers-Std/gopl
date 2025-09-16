// Exercise 4.12: The popular web comic xkcd has a JSON interface. For
// example, a request to https://xkcd.com/571/info.0.json produces a
// detailed description of comic 571, one of many favorites. Download
// each URL (once!) and build an offline index. Write a tool xkcd
// that, using this index, prints the URL and transcript of each comic
// that match es a search term provided on the command line.
package main

import (
	"fmt"
	"log"
	"os"
	"ukiran/gopl/ch4/exercises/ex-4.12/search"
	"ukiran/gopl/ch4/exercises/ex-4.12/xkcd"
)

// mkdir offline
func main() {
	// Fetch and index comics, handle errors
	err := FetchAndIndexAllN(100)
	if err != nil {
		// If there's an error, terminate and print the error
		log.Fatalf("Error fetching and indexing: %v\n", err)
	}

	// Ensure there are keywords passed as arguments
	if len(os.Args) < 2 {
		// If no keywords are provided, exit with a helpful message
		log.Fatal("Please provide at least one search keyword.\n")
	}

	// Extract keywords from command line arguments
	keywords := os.Args[1:]

	// Search for comics matching the keywords
	// This assumes that FetchAndIndexAllN has populated the index correctly
	comics := search.Search(keywords)
	fmt.Println(len(comics))

}

func FetchAndIndexAllN(n int) error {
	for i := 1; i <= n; i++ {
		err := xkcd.FetchAndIndex(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch all n: %v\n", err)
			return err
		}
	}
	return nil
}

// if len(comics) == 0 {
// 	// If no comics are found, inform the user
// 	fmt.Println("No comics found matching the given keywords.")
// } else {
// 	// Print out the comics' URLs and Titles
// 	for _, cp := range comics {
// 		// Print actual values of URL and Title (not pointers)
// 		fmt.Printf("URL: %s, Title: %s\n", cp.URL, cp.Title)
// 	}
// }
