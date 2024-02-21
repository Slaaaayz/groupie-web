package controllers

import (
	"fmt"
	models "groupie-tracker/model"
	"html/template"
	"net/http"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'artiste à partir de l'URL
	parts := strings.Split(r.URL.Path, "/")
	idString := parts[len(parts)-1]

	// Récupère les données de l'artiste
	artist, err := models.GetArtistByID(idString)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}
	artist.DataConcert, err = models.GetData(idString)
	artist.URL = models.UrlSpotify(artist)
	artist.Locations, _ = models.GetLocation(idString)
	if err != nil {
		http.Error(w, "pute", http.StatusInternalServerError)
		return
	}

	// Charger le modèle HTML
	tmpl, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du modèle", http.StatusInternalServerError)
		return
	}
	// Execute
	err = tmpl.Execute(w, artist)
	if err != nil {
		fmt.Println("Erreur d'exécution du modèle:", err)
		fmt.Println(err.Error())
		http.Error(w, "Erreur d'exécution du modèle", http.StatusInternalServerError)
		return
	}
}
