package protocol

import "time"

type GenericRequest struct {
	command string
}

type GenericResponse struct {
	status string 
}

type RequestSites struct {
	GenericRequest
	Type string `json:"type"`
	domain string
}

type ResponseSites struct {
	GenericResponse
	ID int `json:"id"`
	name string 
	page_id int
	lastseen time.Time
}

type PostSites struct {
	GenericRequest
	ID int `json:"id"`
	name string 
	url string 
	page_id int 
	lastseen time.Time
}