package protocol

import "time"

type GenericRequest struct {
	Command string `json:"command"`
}

type GenericResponse struct {
	Status string `json:"status"`
}

type GetSiteRequest struct {
	GenericRequest
	Url string `json:"url"`
}

type GetSiteResponse struct {
	GenericResponse
	Url string `json:"url"`
}

type PostSites struct {
	GenericRequest
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	PageID   int       `json:"page_id"`
	LastSeen time.Time `json:"lastseen"`
}
