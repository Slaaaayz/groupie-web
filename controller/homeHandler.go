package controllers

import (
	models "groupie-tracker/model"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Allartists, err := models.GetArtists()
	var homeArtists []models.Artist
	homeArtists = models.GetRandomArtists(Allartists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("./templates/accueil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, homeArtists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
