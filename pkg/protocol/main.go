package protocol

import "time"

type GenericRequest struct {
	Command string `json:"command"`
}

type GenericResponse struct {
	Status string `json:"status"`
}

type RequestSites struct {
	GenericRequest
	Type   string `json:"type"`
	Domain string `json:"domain"`
}

type ResponseSites struct {
	GenericResponse
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	PageID   int       `json:"page_id"`
	LastSeen time.Time `json:"lastseen"`
}

type PostSites struct {
	GenericRequest
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	PageID   int       `json:"page_id"`
	LastSeen time.Time `json:"lastseen"`
}
