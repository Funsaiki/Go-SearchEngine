package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
)

var sites []donnees.Site

func main() {
	sites = append(sites, donnees.Site{ID: 1, Hostip: "http://5.135.178.104:10987/", Domain: "http://5.135.178.104:10987/", LastSeen: time.Now()})
	sites = append(sites, donnees.Site{ID: 2, Hostip: "http://62.210.124.31/", Domain: "http://62.210.124.31/", LastSeen: time.Now()})

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
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var getSite donnees.Site
		err = json.Unmarshal(body, &getSite)
		if err != nil {
			log.Println("Erreur lors de la conversion des données en structure de demande:", err)
			return
		}

		sites = append(sites, getSite)
		// Définir l'en-tête Content-Type sur application/json
		w.Header().Set("Content-Type", "application/json")

		// Envoyer la réponse
		w.Write([]byte("Site ajouté"))
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
