package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		fmt.Fprintf(w, "item %s is already present\n", item)
		return
	}
	f64, _ := strconv.ParseFloat(price, 32) // ignoring err
	db[item] = dollars(f64)
	fmt.Fprintf(w, "item %s added successfully!\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		fmt.Fprintf(w, "item %s not found\n", item)
		return
	}
	f64, _ := strconv.ParseFloat(price, 32) // ignoring err
	db[item] = dollars(f64)

}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/upadate", http.HandlerFunc(db.update))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
