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
	"io/ioutil"
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
	fmt.Println("Visiting oldest site..." + site.Hostip)
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

	allFiles := getFileRequest()

	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		title := s.Text()
		href, _ := s.Attr("href")
		
		fmt.Printf("Review %d: %s - %s\n", i, title, href)
		
		if len(href) > 1 {
			if href[:2] == "//" {
				href = href[2:]
			} else if href[:1] == "/" {
				href = href[1:]
			}
		}

		var exists bool
		var lastId int

		if allFiles.Files == nil {
			exists = false
		} else {
			last := allFiles.Files[len(allFiles.Files)-1]
			lastId = last.ID + 1
			for _, file := range allFiles.Files {
				if file.Name == title || file.Url == href {
					exists = true
					fmt.Println("File already in database.")
					return 
				} else {
					exists = false
				}
			}
		}

		if !exists {
			// Création de la demande du client
			request := protocol.CreateFileRequest{
				GenericRequest: protocol.GenericRequest{
					Command: "create_file",
				},
				File: donnees.File{ID: lastId + i, Name: title, Url: site.Hostip + href, SiteID: site.ID, LastSeen: time.Now()},
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
			var response protocol.CreateFileResponse
			err = json.Unmarshal(buffer[:n], &response)
			if err != nil {
				log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
			}

			fmt.Println("Response:", response)

			return
		}
	})

	serverAddr := "localhost:5000"

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// Création de la demande du client
	request := protocol.UpdateSiteRequest{
		GenericRequest: protocol.GenericRequest{
			Command: "update_site",
		},
		Site: site,
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

func getFileRequest() protocol.GetFileResponse {
	request := protocol.GetFileRequest{
		GenericRequest: protocol.GenericRequest{
			Command: "get_files",
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
	n, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal("Error receiving response:", err)
	}

	fmt.Println("Received data:", string(n))
	// Conversion des données en structure de réponse
	var response protocol.GetFileResponse
	err = json.Unmarshal(n, &response)
	if err != nil {
		log.Fatal("Erreur lors de la conversion de la réponse en structure de données:", err)
	}

	// Affichage de la réponse
	fmt.Println("Server response:", response)

	return response
}