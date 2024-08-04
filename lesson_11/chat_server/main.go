package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type client chan<- string // an outgoing message channel

type clientInfo struct {
	ch   client
	name string
}

var (
	entering = make(chan clientInfo)
	leaving  = make(chan clientInfo)
	messages = make(chan string) // all incoming client messages
	green    = color.New(color.FgGreen).SprintFunc()
)

var wg sync.WaitGroup

const (
	idleTimeout       = 5 * time.Minute
	messageBufferSize = 100
)

func main() {
	listener := createServer()
	defer listener.Close()

	done := make(chan struct{})

	go broadcaster(done)

	waitForShutdown(listener, done)

	printAddrs()

	handleConnections(listener, done)
}

func createServer() net.Listener {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp", "0.0.0.0:8000", config)
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func printAddrs() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Chat server started on the following local addresses:")
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsPrivate() {
			if ipnet.IP.To4() != nil {
				log.Printf(" - %s:8000\n", ipnet.IP.String())
			}
		}
	}
}

func waitForShutdown(listener net.Listener, done chan<- struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("Received shutdown signal, shutting down gracefully...")
		close(done)
		listener.Close()
		wg.Wait()
		os.Exit(0)
	}()
}

func handleConnections(listener net.Listener, done <-chan struct{}) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			// Check if listener was closed due to graceful shutdown
			select {
			case <-done:
				return
			default:
				log.Print(err)
			}
			continue
		}

		wg.Add(1)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer wg.Done()

	ch := make(chan string, messageBufferSize) // outgoing client messages with buffer

	go clientWriter(conn, ch)

	input, name, who := setNameAndInput(ch, conn)

	waitForMessages(input, conn, name)

	leaving <- clientInfo{ch, name}
	messages <- who + " has left"
	conn.Close()
}

func setNameAndInput(ch chan string, conn net.Conn) (*bufio.Scanner, string, string) {
	ch <- "Enter your name:"
	input := bufio.NewScanner(conn)
	input.Scan()
	name := input.Text()

	who := name + " (" + conn.RemoteAddr().String() + ")"

	ch <- "Welcome, " + name

	messages <- who + " has arrived"
	entering <- clientInfo{ch, name}

	return input, name, who
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			log.Println("Error writing to client:", err)
			return
		}
	}
}

func waitForMessages(input *bufio.Scanner, conn net.Conn, name string) {
	conn.SetReadDeadline(time.Now().Add(idleTimeout)) // Set initial deadline

	for input.Scan() {
		messages <- green("[" + name + "] " + input.Text())

		conn.SetReadDeadline(time.Now().Add(idleTimeout)) // Reset deadline on activity
	}

	if err := input.Err(); err != nil {
		log.Println("Error reading from client:", err)
	}
}

func broadcaster(done chan struct{}) {
	clients := make(map[client]string)

	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients
			for cli := range clients {
				select {
				case cli <- msg:
				default:
					// Drop the message if the client is not ready
				}
			}

		case cliInfo := <-entering:
			clients[cliInfo.ch] = cliInfo.name

			cliInfo.ch <- "Current clients:"
			for _, name := range clients {
				cliInfo.ch <- fmt.Sprintf(" - %s", name)
			}

		case cliInfo := <-leaving:
			delete(clients, cliInfo.ch)
			close(cliInfo.ch)

		case <-done:
			// Graceful shutdown
			for cli := range clients {
				close(cli)
			}
			return
		}
	}
}
