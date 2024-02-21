package controllers

import (
	"html/template"
	"net/http"
)

func BillboardHandler(w http.ResponseWriter, r *http.Request) {
	// Charger le modèle HTML
	tmpl, err := template.ParseFiles("./templates/billboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
