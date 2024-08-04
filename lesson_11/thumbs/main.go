package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"thumb_gen/thumbnail"
)

func main() {
	filenames := make(chan string)

	go func() {
		filenames <- "1.jpg"
		filenames <- "2.jpg"
		filenames <- "3.jpg"
		close(filenames)
	}()

	totalSize := makeThumbnails(filenames)
	fmt.Printf("Total size of thumbnails: %d bytes\n", totalSize)
}

func makeThumbnails(filenames <-chan string) int64 {
	sizes := make(chan int64)

	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go doCreateThumbnail(&wg, f, sizes)
	}

	go doClose(&wg, sizes)

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

func doClose(wg *sync.WaitGroup, sizes chan<- int64) {
	wg.Wait()
	close(sizes)
}

func doCreateThumbnail(wg *sync.WaitGroup, file string, sizes chan<- int64) {
	defer wg.Done()

	thumb, err := thumbnail.ImageFile(file)
	if err != nil {
		log.Println(err)
		return
	}

	info, _ := os.Stat(thumb) // OK to ignore error
	sizes <- info.Size()
}
