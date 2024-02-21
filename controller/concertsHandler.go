package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	models "groupie-tracker/model"
)

func ConcertsHandler(w http.ResponseWriter, r *http.Request) {
	//Récupérer tout les artistes et leurs données des concerts
	allArtists, err := models.GetArtists()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}

	for ind := range allArtists {
		allArtists[ind].DataConcert, _ = models.GetData(strconv.Itoa(allArtists[ind].Id))
	}

	// Charger le modèle HTML
	tmpl, err := template.ParseFiles("./templates/concerts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Execute
	err = tmpl.Execute(w, allArtists)
	if err != nil {
		print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
