package donnees

import "time"

type Site struct {
	ID int `json:"id"`
	Hostip string `json:"hostip"`
	Domain string `json: "domain"`
	LastSeen time.Time `json:"lastseen"`
}

type File struct { 
	ID int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	SiteID int `json:"site_id"`
	LastSeen time.Time `json:"lastseen"`
}

type Database struct {
	Sites  []Site
	Files []File 
}