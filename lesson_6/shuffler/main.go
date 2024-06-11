package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func Shuffle(slice interface{}) {
	rv := reflect.ValueOf(slice)

	if rv.Kind() != reflect.Slice {
		fmt.Println("Shuffle: provided interface is not a slice type")
		return
	}

	length := rv.Len()
	swap := reflect.Swapper(slice)

	rand.Shuffle(length, func(i, j int) {
		swap(i, j)
	})
}

func main() {
	questions := []Question{
		{Country: "France", Capital: "Paris"},
		{Country: "Germany", Capital: "Berlin"},
		{Country: "Italy", Capital: "Rome"},
	}

	fmt.Println("Before shuffle:", questions)
	Shuffle(questions)
	fmt.Println("After shuffle:", questions)

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Before shuffle:", numbers)
	Shuffle(numbers)
	fmt.Println("After shuffle:", numbers)

	Shuffle(42)
}
