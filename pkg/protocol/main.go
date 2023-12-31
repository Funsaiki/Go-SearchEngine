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

type GetFileRequest struct {
	GenericRequest
	Query string `json:"query"`
	Filter string `json:"filter"`
}

type GetFileResponse struct {
	GenericResponse
	Files []donnees.File `json:"files"`
}

type UpdateSiteRequest struct {
	GenericRequest
	Site donnees.Site `json:"site"`
}

type UpdateSiteResponse struct {
	GenericResponse
}

type CreateFileRequest struct {
	GenericRequest
	File donnees.File `json:"file"`
}

type CreateFileResponse struct {
	GenericResponse
	File donnees.File `json:"file"`
}