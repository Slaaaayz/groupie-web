package controllers

import (
	models "groupie-tracker/model"
	"html/template"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var allArtists []models.Artist
	allArtists, _ = models.GetArtists()
	tmpl, err := template.ParseFiles("./templates/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, allArtists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
