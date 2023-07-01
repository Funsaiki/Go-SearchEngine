package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func readInput(out chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		out <- input
	}
}

func receiveData(conn net.Conn, in chan<- string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		in <- line
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	inputCh := make(chan string)
	dataCh := make(chan string)

	go readInput(inputCh)
	go receiveData(conn, dataCh)

	for {
		select {
		case input := <-inputCh:
			fmt.Fprintf(conn, input)
		case data := <-dataCh:
			fmt.Print("Received:", data)
			if strings.Contains(data, "Invalid command") {
				fmt.Println("Please enter a valid command.")
			}
		}
	}
}