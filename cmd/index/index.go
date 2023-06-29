package main

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
//	"time"
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
	var genericRequest protocol.GenericRequest
	err = json.Unmarshal(data, &genericRequest)
	if err != nil {
		log.Println("Erreur lors de la conversion des données en structure de demande:", err)
		return
	}

	fmt.Println("Received command:", genericRequest.Command)

	// Switch sur la commande de demande
	switch genericRequest.Command {
		case "get_sites":
			// Conversion des données en structure de demande GenericRequest
			var request protocol.GetSiteRequest
			err := json.Unmarshal(data, &request)
			if err != nil {
				log.Println("Error decoding create site request:", err)
				return
			}

			// Traitement de la demande et génération de la réponse
			response := protocol.GetSiteResponse{
				GenericResponse: protocol.GenericResponse{
					Status: "ok",
				},
				Url: request.Url,
			}

			// Conversion de la réponse en JSON
			responseData, err := json.Marshal(response)
			if err != nil {
				log.Println("Erreur lors de la conversion de la réponse en JSON:", err)
				return
			}
			fmt.Println("Sending response:", string(responseData))

			// Envoi de la réponse au client
			_, err = conn.Write([]byte(responseData))
			if err != nil {
				log.Println("Error sending create site response:", err)
				return
			}
		default:
			log.Println("Unknown command:", genericRequest.Command)
			return
	}
}