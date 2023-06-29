package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read data from the client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading data:", err)
		return
	}

	// Process the received data
	data := buffer[:n]
	fmt.Println("Received data:", string(data))

	// Send response back to the client
	response := "Hello from server!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println("Error sending response:", err)
		return
	}

	fmt.Println("Response sent:", response)
}