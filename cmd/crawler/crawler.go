package main

import (
	"fmt"
	"log"
	"net"
	"github.com/Funsaiki/Go-SearchEngine/pkg/protocol"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"encoding/json"
	"time"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	for {
		var allSites protocol.GetSiteResponse
		allSites = getSiteRequest()
		var temp, oldest donnees.Site
		var j int
		for i, site := range allSites.Sites {
			fmt.Println("Site:", site.LastSeen)
			fmt.Println("Temp:", temp.LastSeen)
			if temp.LastSeen.IsZero() {
				temp = site
				oldest = temp
				j = i
			} else if site.LastSeen.Before(temp.LastSeen) {
				temp = site
				oldest = temp
				j = i
			}
		}
		fmt.Println("Oldest site:", oldest, "at index", j)
		visitSite(oldest, j)
		time.Sleep(4 * time.Second)
	}
}

func visitSite(site donnees.Site, index int) {
	fmt.Println("Visiting oldest site...")

	// Make HTTP GET request
	res, err := http.Get(site.Hostip)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		href, _ := s.Attr("href")
		if href[:1] == "/" {
			href = href[1:]
		}
		fmt.Printf("Link #%d: %s - %s\n", i, title, site.Hostip + "/" + href)
	})

	// Création de la demande du client
	request := protocol.UpdateSiteRequest{
		GenericRequest: protocol.GenericRequest{
			Command: "update_site",
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
	var response protocol.UpdateSiteResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
	}

	fmt.Println("Response:", response.Status)
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