package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title    string
	Year     int  `json:"released"`
	Color    bool `json:"color,omitempty"`
	Director string
	Actors   []string
}

var movies = []Movie{
	{Title: "Cure", Year: 1997, Color: true,
		Director: "Kiyoshi Kurosawa",
		Actors:   []string{"Koji Yakusho", "Anna Nakagawa"}},
	{Title: "Seven Samurai", Year: 1954, Color: false,
		Director: "Akira Kurosawa",
		Actors:   []string{"Toshiro Mifune", "Takashi Shimura"}},
}

func main() {
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
