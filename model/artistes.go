package models

// Structure qui représente un artiste

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	DataConcert  Data // Contient les dates des concerts et leurs lieux
	URL          string // Url pour l'iframe
	Locations    []string `json:"-"` // Liste des lieux dans lequel l'artiste a un concert
}
