package controllers

import (
	models "groupie-tracker/model"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère un artiste random
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(52)
	randomIndexString := strconv.Itoa(randomIndex)
	randomArtist, err := models.GetArtistByID(randomIndexString)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des artistes", http.StatusInternalServerError)
		return
	}
	// Charger le modèle HTML
	tmpl, err := template.ParseFiles("./templates/game.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute
	err = tmpl.Execute(w, randomArtist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
