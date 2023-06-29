package donnees

import "time"

type Site struct {
	ID int `json:"id"`
	Hostip string `json:"hostip"`
	Domain string `json: "domain"`
	lastseen time.Time
}

type File struct { 
	ID int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	Page string `json:page"`
}

type Database struct {
	Sites  []Site
	Files []File 
}