package main

import (
	"fmt"
	"log"
	"net"
	"encoding/json"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"time"
)

var database donnees.Database

var sites []donnees.Site
var files []donnees.File

func main() {
	sites = append(sites, donnees.Site{ID: 1, Hostip: "http://5.135.178.104:10987/", Domain: "http://5.135.178.104:10987/", LastSeen: time.Now()})
	sites = append(sites, donnees.Site{ID: 2, Hostip: "http://62.210.124.31/", Domain: "http://62.210.124.31/", LastSeen: time.Now()})
	files = append(files, donnees.File{ID: 1, Name: "Donjons%20&%20Dragons%20-%20L%e2%80%99Honneur%20des%20voleurs%20%5bFR-EN%5d%20(2023).mkv", Url: "http://62.210.124.31/Donjons%20&%20Dragons%20-%20L%e2%80%99Honneur%20des%20voleurs%20%5bFR-EN%5d%20(2023).mkv", SiteID: 2, LastSeen: time.Now()})

	database = donnees.Database{
		Sites: sites,
		Files: files,
	}

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening on port 5000...")

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
				Sites: sites,
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
			
			return
		case "create_site":
			fmt.Println("Received command:", genericRequest.Command)
			// Conversion des données en structure de demande GenericRequest
			var request protocol.CreateSiteRequest
			err := json.Unmarshal(data, &request)
			if err != nil {
				log.Println("Error decoding create site request:", err)
				return
			}
			
			sites = append(sites, request.Site)
			fmt.Println("Received site:", request.Site)

			response := protocol.CreateSiteResponse{
				GenericResponse: protocol.GenericResponse{
					Status: "ok",
				},
				Site: request.Site,
			}

			// Conversion de la réponse en JSON
			responseData, err := json.Marshal(response)
			if err != nil {
				log.Println("Erreur lors de la conversion de la réponse en JSON:", err)
				return
			}
			fmt.Println("Sending response:", string(responseData))

			// Traitement de la demande et génération de la réponse
			_, err = conn.Write([]byte(responseData))
			if err != nil {
				log.Println("Error sending create site response:", err)
				return
			}
		case "get_files":
			// Conversion des données en structure de demande GenericRequest
			var request protocol.GetFileRequest
			err := json.Unmarshal(data, &request)
			if err != nil {
				log.Println("Error decoding create site request:", err)
				return
			}

			// Traitement de la demande et génération de la réponse
			response := protocol.GetFileResponse{
				GenericResponse: protocol.GenericResponse{
					Status: "ok",
				},
				Files: files,
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
			
			return
		default:
			log.Println("Unknown command:", genericRequest.Command)
			return
	}
}