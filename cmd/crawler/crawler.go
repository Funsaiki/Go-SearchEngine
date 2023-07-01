package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	serverAddr := "localhost:5000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	message := "Hello from crawler!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal("Error sending data:", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal("Error receiving response:", err)
	}

	response := buffer[:n]
	fmt.Println("Server response:", string(response))
}
