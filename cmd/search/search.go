package main

import (
	"encoding/json"
	"fmt"
	//"io"
	"log"
	"net"
	"net/http"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
)

func main() {
	go http.HandleFunc("/sites", handleSites)

	// En tant que serveur HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from search server!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSites(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		
		// Conversion de la réponse en JSON
		responseData, err := json.Marshal(getSiteRequest())
		if err != nil {
			http.Error(w, "Error encoding TCP response", http.StatusInternalServerError)
			return
		}

		// Définir l'en-tête Content-Type sur application/json
		w.Header().Set("Content-Type", "application/json")

		// Envoyer la réponse
		w.Write(responseData)
	} else if r.Method == http.MethodPost {
		fmt.Println("Received POST request")
		
		// Conversion de la demande en structure de données
		var site donnees.Site
		err := json.NewDecoder(r.Body).Decode(&site)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		// Traitement de la demande et génération de la réponse
		response := createSiteRequest(site)
		
		// Conversion de la réponse en JSON
		responseData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error encoding TCP response", http.StatusInternalServerError)
			return
		}
		// Définir l'en-tête Content-Type sur application/json
		w.Header().Set("Content-Type", "application/json")

		// Envoyer la réponse
		w.Write(responseData)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func getSiteRequest() protocol.GetSiteResponse {
	// Création de la demande du client
	request := protocol.GetSiteRequest{
		GenericRequest: protocol.GenericRequest{
			Command: "get_sites",
		},
	}

	serverAddr := "localhost:5000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

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
	var response protocol.GetSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
	}

	// Affichage de la réponse
	fmt.Println("Server response:", response)

	return response
}

func createSiteRequest(site donnees.Site) protocol.CreateSiteResponse {
	// Création de la demande du client
	request := protocol.CreateSiteRequest{
		GenericRequest: protocol.GenericRequest{
			Command: "create_site",
		},
		Site: site,
	}

	serverAddr := "localhost:5000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

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
	var response protocol.CreateSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
	}

	// Affichage de la réponse
	fmt.Println("Server response:", response)

	return response
}