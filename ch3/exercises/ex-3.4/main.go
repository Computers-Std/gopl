package main

import (
	"log"
	"net/http"
	"ukiran/gopl/ch3/exercises/ex-3.4/surface"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w, 600, 320)
}
