package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
//	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"time"
)

type Site struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	PageID   int       `json:"page_id"`
	LastSeen time.Time `json:"lastseen"`
}

var sites []Site

func main() {
	sites = append(sites, Site{ID: 1, Name: "Site 1", URL: "https://site1.com", PageID: 123, LastSeen: time.Now()})
	sites = append(sites, Site{ID: 2, Name: "Site 2", URL: "https://site2.com", PageID: 456, LastSeen: time.Now()})

	http.HandleFunc("/sites", handleSites)
	
	// En tant que client
	go func() {
		serverAddr := "localhost:8080"

		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			log.Fatal("Error connecting to server:", err)
		}
		defer conn.Close()

		// Création de la demande du client
		request := protocol.GetSiteRequest{
			GenericRequest: protocol.GenericRequest{
				Command: "get_sites",
			},
			Url: "http://62.210.124.31/",
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
		var response protocol.GetSiteResponse
		err = json.Unmarshal(buffer[:n], &response)
		if err != nil {
			log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
		}
		
		// Affichage de la réponse
		fmt.Println("Server response:", response)
	}()

	// En tant que serveur HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from search server!")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

	url := "http://localhost:8081/sites"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var sites []string
	err = json.NewDecoder(response.Body).Decode(&sites)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sites)
}

func handleSites(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Récupérer la liste des sites
		responseData, err := json.Marshal(sites)
		if err != nil {
			http.Error(w, "Error encoding sites response", http.StatusInternalServerError)
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