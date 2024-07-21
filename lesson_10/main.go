package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squarer(naturals, squares)

	printer(squares)
}

func counter(naturals chan<- int) {
	for x := 0; x < 100; x++ {
		naturals <- x
	}
	close(naturals)
}

func squarer(naturals <-chan int, squares chan<- int) {
	// for x := range naturals {
	// 	squares <- x * x
	// }
	// close(squares)

	for {
		select {
		case x, ok := <-naturals:
			if !ok {
				fmt.Println("Naturals is closed")
				// If the channel is closed, break out of the loop
				close(squares)
				return
			}
			squares <- x * x
		case <-time.After(time.Second):
			fmt.Println("Timeout: no data received from naturals")
			close(squares)
			return
		}
	}
}

func printer(squares <-chan int) {
	// for x := range squares {
	// 	fmt.Println(x)
	// }

	for {
		select {
		case x, ok := <-squares:
			if !ok {
				fmt.Println("Squares is closed")
				// If the channel is closed, break out of the loop
				return
			}
			fmt.Println(x)
		case <-time.After(time.Second):
			fmt.Println("Timeout: no data received from squares")
			return
		}
	}
}

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		<-ch
// 	}()

// 	ch <- 42
// }

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	done := make(chan bool)

// 	go spinner(100*time.Millisecond, done)

// 	work()
// 	done <- true
// }

// func work() {
// 	time.Sleep(5 * time.Second)
// }

// func spinner(delay time.Duration, done chan bool) {
// 	for {
// 		select {
// 		case <-done:
// 			fmt.Println("done")
// 			return
// 		default:
// 			for _, r := range `-\|/` {
// 				fmt.Printf("\r%c", r)
// 				time.Sleep(delay)
// 			}
// 		}
// 	}
// }

// import (
// 	"fmt"
// 	"math/rand"
// 	"runtime"
// 	"time"
// )

// func request(server string) string {
// 	delay := rand.Intn(1000)

// 	fmt.Printf("Requesting %s (will take %d ms)\n", server, delay)
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	fmt.Printf("Response received from %s\n", server)

// 	return fmt.Sprintf("Response from %s", server)
// }

// func mirroredQuery() string {
// 	responses := make(chan string, 3)

// 	go func() { responses <- request("asia.gopl.io") }()
// 	go func() { responses <- request("europe.gopl.io") }()
// 	go func() { responses <- request("americas.gopl.io") }()

// 	return <-responses
// }

// func main() {
// 	result := mirroredQuery()
// 	fmt.Println("Final result in main", result)

// 	// Wait for a while to see if any goroutines are still running
// 	time.Sleep(2 * time.Second)

// 	// Display the number of currently running goroutines
// 	fmt.Printf("Number of goroutines after main: %d\n", runtime.NumGoroutine())
// }

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Server is listening on localhost:8000")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Print(err) // e.g., connection aborted
// 			continue
// 		}
// 		log.Printf("Connection established with %s", conn.RemoteAddr())
// 		go handleConn(conn) // handle each connection concurrently
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer func() {
// 		log.Printf("Connection closed with %s", c.RemoteAddr())
// 		c.Close()
// 	}()

// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			log.Printf("Error writing to %s: %v", c.RemoteAddr(), err)
// 			return // e.g., client disconnected
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

// const (
// 	SPIN_CHARS = `-\|/`
// )

// func main() {
// 	go spinner(100 * time.Millisecond)

// 	fmt.Println("Main function doing work")
// 	go work()
// 	fmt.Println("Main function done")

// 	time.Sleep(6 * time.Second)
// }

// func spinner(delay time.Duration) {
// 	fmt.Println("Starting go routine")

// 	for index := 0; ; index = (index + 1) % len(SPIN_CHARS) {
// 		fmt.Printf("\r%c", SPIN_CHARS[index])
// 		time.Sleep(delay)
// 	}
// }

// func work() {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("WORK done")
// }
