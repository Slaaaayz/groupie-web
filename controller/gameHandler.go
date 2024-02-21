package controllers

import (
	models "groupie-tracker/model"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {

	listArtists, err1 := models.GetArtists()
	if err1 != nil {
		print(err1.Error())
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(listArtists))

	tmpl, err := template.ParseFiles("./templates/game.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, listArtists[randomIndex])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
