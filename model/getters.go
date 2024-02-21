package models

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func GetData(id string) (Data, error) { // Récupère une data de type Data à partir de l'api relation en fonction de l'id de l'artiste choisi
	var data Data
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	// Utiliser une structure intermédiaire pour décoder la réponse JSON
	var intermediateStruct struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	err = json.NewDecoder(resp.Body).Decode(&intermediateStruct)
	if err != nil {
		return data, err
	}

	// Copier les valeurs de la structure intermédiaire vers la structure finale
	data.Id = intermediateStruct.Id
	data.DatesLocations = intermediateStruct.DatesLocations
	for k, _ := range data.DatesLocations {
		data.DatesLocations[format(k)] = data.DatesLocations[k]
		delete(data.DatesLocations, k)
	}
	return data, nil
}

func GetLocation(id string) ([]string, error) { // Récupère une liste de string correspondant au lieux de l'artiste à partir de l'API location en fonction de l'artiste choisi
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Utiliser une structure intermédiaire pour décoder la réponse JSON
	var intermediateStruct struct {
		Locations []string `json:"locations"`
	}

	err = json.NewDecoder(resp.Body).Decode(&intermediateStruct)
	if err != nil {
		return intermediateStruct.Locations, err
	}
	return intermediateStruct.Locations, nil
}

func GetRandomArtists(allArtists []Artist) []Artist { //Choisis 8 artistes au hasard parmi tout les artistes (pour la page d'accueil)
	rand.Seed(time.Now().UnixNano())
	artistsCopy := make([]Artist, len(allArtists))
	copy(artistsCopy, allArtists)
	randomArtists := make([]Artist, 0, 8)
	for i := 0; i < 8; i++ {
		randomIndex := rand.Intn(len(artistsCopy))
		randomArtists = append(randomArtists, artistsCopy[randomIndex])
		artistsCopy = append(artistsCopy[:randomIndex], artistsCopy[randomIndex+1:]...)
	}
	return randomArtists
}

func GetArtists() ([]Artist, error) { // Récupère tout les artistes dans une liste de structure de type Artiste
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		print(err.Error())
		return nil, err
	}
	return artists, nil
}

func GetArtistByID(id string) (Artist, error) { // Récupère tout les artistes dans une liste de structure de type Artiste
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	var artists Artist
	if err != nil {
		return artists, err
	}
	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		print(err.Error())
		return artists, err
	}
	return artists, nil
}


func format(key string) string {
	var formattedKey string
	capitalizeNext := true
	for _, char := range key {
		if char == '-' {
			break
		} else if char == '_' {
			formattedKey += " "
			capitalizeNext = true
		} else {
			if capitalizeNext {
				formattedKey += strings.ToUpper(string(char))
				capitalizeNext = false
			} else {
				formattedKey += string(char)
			}
		}
	}
	return formattedKey
}
