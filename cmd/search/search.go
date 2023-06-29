package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	// En tant que client
	go func() {
		serverAddr := "localhost:8080"

		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			log.Fatal("Error connecting to server:", err)
		}
		defer conn.Close()

		message := "Hello from search client!"
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
	}()

	// En tant que serveur HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from search server!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
