package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// flag
	fmt.Println("Enter hex number or 'stop' to exit:")

	var input string

	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		if input == "stop" {
			break
		}

		i := new(big.Int)

		if _, ok := i.SetString(processHex(input), 16); !ok {
			fmt.Println("Invalid hexadecimal number!")
			continue
		}

		fmt.Println(i)
	}
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
