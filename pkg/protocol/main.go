package protocol

import (
	"time"
	"github.com/Funsaiki/Go-SearchEngine/pkg/donnees"
)

type GenericRequest struct {
	Command string `json:"command"`
}

type GenericResponse struct {
	Status string `json:"status"`
}

type GetSiteRequest struct {
	GenericRequest
	Query string `json:"query"`
	Filter string `json:"filter"`
}

type GetSiteResponse struct {
	GenericResponse
	Sites []donnees.Site `json:"sites"`
}

type PostSites struct {
	GenericRequest
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	URL      string    `json:"url"`
	PageID   int       `json:"page_id"`
	LastSeen time.Time `json:"lastseen"`
}
