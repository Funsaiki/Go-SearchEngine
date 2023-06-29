package main

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
	"time"
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

	// Conversion des données en structure de demande
	var request protocol.RequestSites
	err = json.Unmarshal(data, &request)
	if err != nil {
		log.Println("Erreur lors de la conversion des données en structure de demande:", err)
		return
	}

	// Traitement de la demande et génération de la réponse
	response := protocol.ResponseSites{
		GenericResponse: protocol.GenericResponse{
			Status: "success",
		},
		ID:       1,
		Name:     "example",
		PageID:   123,
		LastSeen: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	// Conversion de la réponse en JSON
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println("Erreur lors de la conversion de la réponse en JSON:", err)
		return
	}

	// Envoi de la réponse au client
	_, err = conn.Write(responseData)
	if err != nil {
		log.Println("Error sending response:", err)
		return
	}

	fmt.Println("Response sent:", response)
}