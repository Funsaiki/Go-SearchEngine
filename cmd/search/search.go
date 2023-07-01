package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"log"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
	"time"
)

var sites []donnees.Site

func readInput(out chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		out <- input
	}
}

func receiveData(conn net.Conn, in chan<- string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		in <- line
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request, inputCh chan<- string) {
	command := r.URL.Path[1:]
	inputCh <- command
}

func main() {
	sites = append(sites, donnees.Site{ID: 1, Hostip: "donnees.Site 1", Domain: "https://site1.com", LastSeen: time.Now()})
	sites = append(sites, donnees.Site{ID: 2, Hostip: "donnees.Site 2", Domain: "https://site2.com", LastSeen: time.Now()})

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	inputCh := make(chan string)
	dataCh := make(chan string)

	go readInput(inputCh)
	go receiveData(conn, dataCh)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from search server!")
	})

	go func() {
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
	}()

	for {
		select {
		case input := <-inputCh:
			fmt.Fprintf(conn, input)
		case data := <-dataCh:
			fmt.Print("Received:", data)
			if strings.Contains(data, "Invalid command") {
				fmt.Println("Please enter a valid command.")
			}
		}
	}
}