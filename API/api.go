package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func artists() []Artist {
	var a []Artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &a)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return a
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		artists := artists()
		for _, artist := range artists {
			fmt.Fprintf(w, "Image: %s\n", artist.Image)
			fmt.Fprintf(w, "Nom: %s\n", artist.Name)
			fmt.Fprintf(w, "Membres: %s\n", artist.Members)
			fmt.Fprintf(w, "Date of creation: %s\n", artist.CreationDate)
			fmt.Fprintf(w, "Date of first Album: %s\n", artist.FirstAlbum)
			fmt.Fprintf(w, "Location: %s\n", artist.Locations)
			fmt.Fprintf(w, "Date of concert: %s\n", artist.ConcertDates)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

