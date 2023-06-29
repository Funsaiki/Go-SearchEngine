package donnees

import "time"

type Dir struct {
	ID int `json:"id"`
	Hostname string `json:"hostname"`
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
	Dirs  []Dir  
	Files []File 
}