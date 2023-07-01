package protocol

import (
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

type CreateSiteRequest struct {
	GenericRequest
	Site donnees.Site `json:"site"`
}

type CreateSiteResponse struct {
	GenericResponse
	Site donnees.Site `json:"site"`
}