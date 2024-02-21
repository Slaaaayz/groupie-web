package models

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

// Structure pour parser la réponse JSON contenant les top tracks
type TopTracksResponse struct {
    Tracks []struct {
        Name string `json:"name"`
    }   `json:"tracks"`
}

func main() {
    // Votre token d'accès
    accessToken := "ead104de30824942869ab90853f5816a"
    // L'ID de l'artiste
    artistID := "0k17h0D3J5VfsdmQ1iZtE9"
    // Le marché (pays)
    market := "FR"

    // Construire la requête
    url := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/top-tracks?market=%s", artistID, market)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Erreur lors de la création de la requête :", err)
        return
    }

    // Ajouter le header d'authentification
    req.Header.Add("Authorization", "Bearer "+accessToken)

    // Exécuter la requête
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Erreur lors de l'exécution de la requête :", err)
        return
    }
    defer resp.Body.Close()

    // Lire la réponse
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Erreur lors de la lecture de la réponse :", err)
        return
    }

    // Parser la réponse JSON
    var topTracks TopTracksResponse
    err = json.Unmarshal(body, &topTracks)
    if err != nil {
        fmt.Println("Erreur lors du parsing de la réponse JSON :", err)
        return
    }

    // Afficher la piste la plus écoutée
    if len(topTracks.Tracks) > 0 {
        fmt.Printf("La musique la plus écoutée de l'artiste : %s\n", topTracks.Tracks[0].Name)
    } else {
        fmt.Println("Aucune piste trouvée pour cet artiste.")
    }
}