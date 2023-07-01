package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"time"
	"io"
)

var sites []donnees.Site

func main() {
	sites = append(sites, donnees.Site{ID: 1, Hostip: "http://5.135.178.104:10987/", Domain: "http://5.135.178.104:10987/", LastSeen: time.Now()})
	sites = append(sites, donnees.Site{ID: 2, Hostip: "http://62.210.124.31/", Domain: "http://62.210.124.31/", LastSeen: time.Now()})
	
	// En tant que client
	go func() {
		serverAddr := "localhost:5000"

		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			log.Fatal("Error connecting to server:", err)
		}
		defer conn.Close()

		http.HandleFunc("/sites", handleSites)

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

	log.Fatal(http.ListenAndServe(":8080", nil))

	url := "http://localhost:8080/sites"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&sites)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sites)
}

func handleSites(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sites", sites)
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
		w.Write([]byte("test"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}