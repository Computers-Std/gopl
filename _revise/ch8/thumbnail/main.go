package main

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

// makeThumbnails makes thumbnails for the specified files
func makeThumbnails1(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// NOTE incorrect!
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

// makeThumbnails3 makes thumbnails of specified files in parallel
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		// NOTE passing "f" as an explicit arg
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE ignoring errors
			ch <- struct{}{}
		}(f)
	}
	// Wait for goroutines to complete
	for range filenames {
		<-ch
	}
}

// makeThumbnails4 make thumbnails for the specified files in
// parallel. It returns an error if any step failed.
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE incorrect: goroutine leak!
		}
	}
	return nil
}

// makeThumbnails5 makes thumbnails for the specified files in
// parallel. It returns the generated file names in an arbitrary
// order, or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	// If an error is encountered in the main goroutine loop, and the
	// function executes "return nil, it.err", the preceding
	// goroutines will generally continue to run and populate the
	// channel "ch". They do not automatically terminate. Simply,
	// because the function that launched them has "retured".
	return thumbfiles, nil
}

// makeThumbnails6 makes thumbnails for each file recieved from the
// channel. It returns the number of bytes occupied by the files it
// creates.
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	// main
	var total int64
	// this loop will not terminate, if the sizes channel is not
	// closed.
	for size := range sizes {
		total += size
	}
	// chech "Figure 8.5"
	return total
}
