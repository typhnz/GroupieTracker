package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"text/template"
)

type ArtistAPI struct {
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

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"DatesLocations"`
}

type ArtistsArray struct {
	Artists []ArtistAPI
}

type Description struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

// 3
var tempCard = template.Must(template.ParseFiles("HTML/artists.html"))    //artists -> card
var tempHome = template.Must(template.ParseFiles("templates/HTML/home.html"))      //hpage -> homePage
var tempDetails = template.Must(template.ParseFiles("HTML/details.html")) //details -> details
var apiElements []ArtistAPI

func artist(w http.ResponseWriter, r *http.Request) {

	api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	apiData, err := ioutil.ReadAll(api.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(apiData, &apiElements)

	artistsData := ArtistsArray{
		Artists: apiElements,
	}

	artistsMap := map[string]interface{}{
		"DataArtists": artistsData,
	}
	tempCard.Execute(w, artistsMap)
}

func home(w http.ResponseWriter, r *http.Request) {

	// api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	// var apiElements ArtistAPI

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// apiDataArtist, err := ioutil.ReadAll(api.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// json.Unmarshal(apiDataArtist, &apiElements)

	tempHome.Execute(w, "home" /*err*/)
}

// 8
func details(w http.ResponseWriter, r *http.Request) {

	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	var locationsObject Relation

	dataArtists := Description{
		ID:           apiElements[pathIDint-1].ID,
		Image:        apiElements[pathIDint-1].Image,
		Members:      apiElements[pathIDint-1].Members,
		CreationDate: apiElements[pathIDint-1].CreationDate,
		FirstAlbum:   apiElements[pathIDint-1].FirstAlbum,
		Locations:    apiElements[pathIDint-1].Locations,
		ConcertDates: apiElements[pathIDint-1].ConcertDates,
		Relations:    apiElements[pathIDint-1].Relations,
	}

	relations, err := http.Get(dataArtists.Relations)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	relationsData, err := ioutil.ReadAll(relations.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(relationsData, &locationsObject)

	mapInt := map[string]interface{}{
		"DataArtists": dataArtists,
		"Relations":   locationsObject,
	}

	tempDetails.Execute(w, mapInt)
}

// 9
func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/artist/", details)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

