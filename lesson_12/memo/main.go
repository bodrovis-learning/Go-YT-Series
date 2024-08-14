package main

import (
	"fmt"
	"io"
	"log"
	"memo/memo"
	"net/http"
	"sync"
	"time"
)

func main() {
	m := memo.New(httpGetBody)
	var n sync.WaitGroup
	urls := []string{
		"https://golang.org",
		"https://pkg.go.dev",
		"https://go.dev/doc/",
		"https://pkg.go.dev",
		"https://blog.golang.org",
		"https://tour.golang.org",
		"https://golang.org",
	}

	for _, url := range urls {
		n.Add(1)
		time.Sleep(1 * time.Second)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
