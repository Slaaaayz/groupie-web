package controllers

import (
	models "groupie-tracker/model"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//Récupérer 8 artistes random pour la page d'accueil
	Allartists, err := models.GetArtists()
	homeArtists := models.GetRandomArtists(Allartists)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	// Charger le modèle HTML
	tmpl, err := template.ParseFiles("./templates/accueil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute
	err = tmpl.Execute(w, homeArtists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
