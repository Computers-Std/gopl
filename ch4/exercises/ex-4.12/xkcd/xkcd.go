package xkcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Comic struct {
	Number   int    `json:"num"`
	Year     string `json:"year"`
	Title    string `json:"title"`
	AltTitle string `json:"alt"`
	URL      string
}

const URL = "https://xkcd.com/"
const FILE_DIR = "offline"
const FILE_EXT = ".json"

var ComicsIndex = make(map[int]Comic)

// download and index nth comic
func FetchAndIndex(n int) error {
	indexStr := strconv.Itoa(n)
	urlPath := URL + indexStr + "/info.0.json"
	filePath := path.Join(FILE_DIR, indexStr+FILE_EXT)
	var comic Comic

	if _, ok := ComicsIndex[n]; !ok {
		if data, err := os.Open(filePath); errors.Is(err, os.ErrNotExist) {

			// fetch file and index
			body, err := fetchURL(urlPath)
			if err != nil {
				return fmt.Errorf("fetchURL: %v\n", err)
			}
			err = IndexData(&comic, body)
			if err != nil {
				return fmt.Errorf("fetch: IndexData: %v", err)
			}
			comic.URL = urlPath
			ComicsIndex[n] = comic

			// create file
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "creating file: %v\n", err)
			}
			defer file.Close() // close the file when done

			// write the response to file
			_, err = io.Copy(file, bytes.NewReader(body))
			if err != nil {
				fmt.Fprintf(os.Stderr, "writing to file: %v\n", err)
			}
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "file exist but, a reading error: %v\n", err)
		} else {
			// index the local file
			body, err := io.ReadAll(data)
			if err != nil {
				return fmt.Errorf("file exist but, error reading: %v\n", err)
			}
			err = IndexData(&comic, body)
			if err != nil {
				return fmt.Errorf("file exist but, error indexing: %v\n", err)
			}
			comic.URL = urlPath
			ComicsIndex[n] = comic
		}

	}
	return nil
}

func IndexData(comic *Comic, body []byte) error {
	if err := json.Unmarshal(body, &comic); err != nil {
		return fmt.Errorf("unmarshal response: %v\n", err)
	}
	return nil
}

func fetchURL(url string) ([]byte, error) {
	// http.Get + write to file + json.Unmarshal to map
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching: %v", err)
	}
	defer resp.Body.Close() // Ensure we close the response body

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body into memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetch: reading response body: %v", err)
	}
	return body, nil
}
