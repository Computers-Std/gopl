package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"ukiran/gopl/ch1/exercises/ex-1.12/lissajous"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := strings.SplitAfter(r.URL.RawQuery, "=")
	cycles, _ := strconv.Atoi(query[1])
	lissajous.Lissajous(w, cycles)
}
