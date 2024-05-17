package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var results []int

	fmt.Print("a\\b\t")

	for _, a := range numbers {
		fmt.Printf("%8d", a)
	}

	fmt.Print("\n\n")

	for _, a := range numbers {
		fmt.Printf("%d\t", a)

		for _, b := range numbers {
			result := a * b
			fmt.Printf("%8d", result)

			results = append(results, result)
		}

		fmt.Println()
	}

	fmt.Println("\n\tResults")
	for index, result := range results {
		fmt.Printf("%8d\t%8d\n", index, result)
	}
}
