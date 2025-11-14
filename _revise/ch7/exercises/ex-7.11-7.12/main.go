package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

// create
func (db database) create(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("item") || !r.URL.Query().Has("price") {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid request, create expects item&price")
		return
	}

	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "item already in the list: %s\n", item)
		return
	}
	f64, _ := strconv.ParseFloat(price, 32) // ignoring err
	db[item] = dollars(f64)
	fmt.Fprintf(w, "added successfully! %v: %v\n", item, price)
}

var index = template.Must(template.ParseFiles("table.tmpl"))

// read
func (db database) read(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("item") && !r.URL.Query().Has("price") {
		// default read request, prints list
		// db.list(w, r)
		if err := index.Execute(w, db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if r.URL.Query().Has("item") {
		// print the item's price
		db.price(w, r)
		return
	}
}

// update
func (db database) update(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("item") || !r.URL.Query().Has("price") {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid request, update expects item&price")
		return
	}
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not in the list: %s\n", item)
		return
	}
	f64, _ := strconv.ParseFloat(price, 32) // ignoring err
	db[item] = dollars(f64)
	fmt.Fprintf(w, "updated successfully! %v: %v\n", item, price)
}

// delete
func (db database) delete(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("item") {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid request, delete requests item")
		return
	}
	item := r.URL.Query().Get("item")
	delete(db, item)
	fmt.Fprintf(w, "item deleted successfully %s\n", item)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/read", db.read)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
