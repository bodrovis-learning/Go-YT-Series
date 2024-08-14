package main

import (
	"fmt"
	"map_data/map_data"
	"sync"
)

var loadDataOnce sync.Once
var allData map[string]int

func main() {
	var wg sync.WaitGroup

	names := []string{"a", "bb", "ccc", "dddd"}

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				fmt.Printf("Goroutine for %s: Icon value = %d\n", name, Data(name))
			}
		}(name)
	}

	wg.Wait()
}

func Data(name string) int {
	loadDataOnce.Do(loadData)

	return allData[name]
}

func loadData() {
	allData = map[string]int{
		"a":    map_data.LoadData("a"),
		"bb":   map_data.LoadData("bb"),
		"ccc":  map_data.LoadData("ccc"),
		"dddd": map_data.LoadData("dddd"),
	}
}
