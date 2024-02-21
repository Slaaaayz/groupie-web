package main

import (
	"fmt"
	//"io/ioutil"
	controllers "groupie-tracker/controller"
	"net/http"
)

// Lancer le serveur

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/artist/", controllers.ArtistHandler)
	http.HandleFunc("/search", controllers.SearchHandler)
	http.HandleFunc("/concerts", controllers.ConcertsHandler)
	http.HandleFunc("/billboard", controllers.BillboardHandler)
	http.HandleFunc("/settings", controllers.SettingsHandler)
	http.HandleFunc("/game", controllers.PlayHandler)
	http.HandleFunc("/result", controllers.ResultHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
