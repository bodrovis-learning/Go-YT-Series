package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define command-line arguments
	serverAddr := flag.String("server", "localhost:8000", "server address in the format ip:port")
	flag.Parse()

	conn, err := makeConn(*serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", *serverAddr)

	// Channel to signal when the server connection is closed
	done := make(chan struct{})

	// Start a goroutine to read messages from the server and print them
	go readMessages(conn, done)

	// Read input from the user and send it to the server
	go sendMessage(conn, done)

	// Wait for the done signal indicating the server connection is closed
	<-done
	fmt.Println("Disconnected from server.")
}

// makeConn establishes a TLS connection to the server.
func makeConn(serverAddr string) (*tls.Conn, error) {
	conf := &tls.Config{
		InsecureSkipVerify: true, // Insecure, should be used with caution
	}

	return tls.Dial("tcp", serverAddr, conf)
}

func readMessages(conn *tls.Conn, done chan<- struct{}) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from server:", err)
	}
	close(done)
}

func sendMessage(conn *tls.Conn, done chan<- struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "/exit" {
			fmt.Println("Exiting chat...")
			conn.Close()
			close(done)
			return
		}

		_, err := fmt.Fprintln(conn, text)
		if err != nil {
			log.Println("Error sending to server:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from stdin:", err)
	}
}
