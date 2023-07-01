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

type CreateSiteRequest struct {
	GenericRequest
	Url string `json:"url"`
}

type CreateSiteResponse struct {
	GenericResponse
	Url string `json:"url"`
}

type GetFileRequest struct {
	GenericRequest
	Url string `json:"url"`
}

type GetFileResponse struct {
	GenericResponse
	Url string `json:"url"`
}

type CreateFileRequest struct {
	GenericRequest
	Url string `json:"url"`
}

type CreateFileResponse struct {
	GenericResponse
	Url string `json:"url"`
}

type UpdateSiteRequest struct {
	GenericRequest
	Url string `json:"url"`
}

type UpdateSiteResponse struct {
	GenericResponse
	Url string `json:"url"`
}

type UpdateFileRequest struct {
	GenericRequest
	Url string `json:"url"`
}