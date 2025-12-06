package main

import (
	"html/template"
	"net/http"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var index = template.Must(template.ParseFiles("table.tmpl"))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func handler(w http.ResponseWriter, r *http.Request) {
	field := r.URL.Path[1:]
	switch field {
	case "title":
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			return x.Title < y.Title
		}})
	case "artist":
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			return x.Artist < y.Artist
		}})
	case "album":
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			return x.Album < y.Album
		}})
	case "year":
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			return x.Year < y.Year
		}})
	case "length":
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			return x.Length < y.Length
		}})
	}
	if err := index.Execute(w, tracks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
