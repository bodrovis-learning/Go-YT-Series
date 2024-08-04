package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		var b = make([]byte, 1)

		for {
			os.Stdin.Read(b)

			if b[0] == 'Q' || b[0] == 'q' {
				abort <- struct{}{}
				return
			}
		}
	}()

	fmt.Println("Commencing countdown. Press 'Q' to abort.")

	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			fmt.Println("...")
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
