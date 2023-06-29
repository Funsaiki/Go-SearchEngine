package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
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

		// Création de la demande du client
		request := protocol.RequestSites{
			GenericRequest: protocol.GenericRequest{
				Command: "get_sites",
			},
			Type:   "example",
			Domain: "example.com",
		}

		// Conversion de la demande en JSON
		requestData, err := json.Marshal(request)
		if err != nil {
			log.Fatal("Erreur lors de la conversion de la demande en JSON:", err)
		}

		// Envoi de la demande via la connexion TCP
		_, err = conn.Write(requestData)
		if err != nil {
			log.Fatal("Erreur lors de l'envoi de la demande:", err)
		}

		// Lecture de la réponse du serveur depuis la connexion TCP
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal("Error receiving response:", err)
		}

		// Conversion des données en structure de réponse
		var response protocol.ResponseSites
		err = json.Unmarshal(buffer[:n], &response)
		if err != nil {
			log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
		}
	}()

	// En tant que serveur HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from search server!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
