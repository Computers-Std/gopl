package search

import (
	"strings"
	"ukiran/gopl/ch4/exercises/ex-4.12/xkcd"
)

//	func Search(keywords []string) []*xkcd.Comic {
//		var result []*xkcd.Comic
//		for _, comic := range xkcd.ComicsIndex {
//			// Check if any keyword is in the Title or AltTitle
//			for _, word := range keywords {
//				if strings.Contains(comic.Title, word) || strings.Contains(comic.AltTitle, word) {
//					result = append(result, &comic)
//					// break
//				}
//			}
//		}
//		return result
//	}
var result []*xkcd.Comic

func Search(keywords []string) []*xkcd.Comic {
	for _, comic := range xkcd.ComicsIndex {
		// Check if any keyword is in the Title or AltTitle
		for _, word := range keywords {
			if strings.Contains(comic.Title, word) || strings.Contains(comic.AltTitle, word) {
				// Make a copy of the comic and append its reference to result
				copyComic := comic
				result = append(result, &copyComic)
				// break // Stop after the first keyword match for this comic
			}
		}
	}
	return result
}
