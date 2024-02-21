package models

type Data struct {
	Id             int                 `json:"id"` // id de l'artiste
	DatesLocations map[string][]string `json:"datesLocations"` // keys = locations, values = liste de dates
}
