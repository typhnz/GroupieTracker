package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

<<<<<<< HEAD
type Artist struct {
=======
type artist struct {
>>>>>>> c705692758414836dff95ee2ffb94cc7a5ad7a09
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
<<<<<<< HEAD
	Relations    string   `json:"relations"`
}
=======
	//Relation []string
}

var element []artist
>>>>>>> c705692758414836dff95ee2ffb94cc7a5ad7a09

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
<<<<<<< HEAD
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

=======
	for i := 0; i < len(a); i++ {
		fmt.Println("Image:", a[i].Image)
		fmt.Println("Name:", a[i].Name)
		fmt.Println("Members:", a[i].Members)
		fmt.Println("CreationDate:", a[i].CreationDate)
		fmt.Println("FirstAlbum", a[i].FirstAlbum)
		fmt.Println("Locations:", a[i].Locations)
		fmt.Println("ConcertDates:", a[i].ConcertDates)
		fmt.Printf("\n")
	}
}

func main() {
	artists()
}
>>>>>>> c705692758414836dff95ee2ffb94cc7a5ad7a09
