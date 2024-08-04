package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var done = make(chan struct{})
var wg sync.WaitGroup

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	fmt.Println("Press enter to cancel the tasks.")

	completion := make(chan int)

	for i := 0; i < 5; i++ {
		if cancelled() {
			fmt.Println("Main loop cancelled!")
			break
		}

		wg.Add(1)
		go longRunningTask(i, completion)
	}

	go func() {
		wg.Wait()
		close(completion)
	}()

	for result := range completion {
		fmt.Printf("Task %d completed\n", result)
	}

	fmt.Println("Main function completed.")
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func longRunningTask(id int, completion chan<- int) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		select {
		case <-done:
			fmt.Printf("Task %d cancelled during execution!\n", id)
			return
		default:
			// Simulate some work with a sleep
			fmt.Printf("Task %d working... %d\n", id, i)
			time.Sleep(1 * time.Second)
		}
	}

	completion <- id
	fmt.Printf("Task %d completed!\n", id)
}
