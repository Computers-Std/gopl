// NOTE: mkdir posters
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const API_KEY = "[YOUR_KEY]"
const URL = "http://www.omdbapi.com/?"
const POSTER_DIR = "posters"

type OMDB struct {
	Title  string
	Year   string
	Poster string
	ImdbID string `json:"imdbID"`
}

// fetch movie poster
func Fetch(id string) (OMDB, error) {
	resp, err := http.Get(URL + "apikey=" + API_KEY + "&i=" + id)
	if err != nil {
		return OMDB{}, fmt.Errorf("fetching: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return OMDB{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result OMDB
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return OMDB{}, fmt.Errorf("Decode error: %v\n", err)
	}
	return result, nil
}

func getPoster(film OMDB) error {
	resp, err := http.Get(film.Poster)
	if err != nil {
		return fmt.Errorf("poster: %v\n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	filePath := path.Join(POSTER_DIR, film.ImdbID+".jpg")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file: %v\n", err)
	}
	defer file.Close()

	// write the response to file
	_, err = io.Copy(file, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("writing to file: %v\n", err)
	}
	return nil
}

func main() {
	arg := os.Args[1]
	film, err := Fetch(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	err = getPoster(film)
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
	}
}
