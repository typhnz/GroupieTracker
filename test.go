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

type Descritpion struct {
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

//3
var tempCard = template.Must(template.ParseFiles("HTML/artists.html"))    //templates -> tempCard           check      artists -> card
var tempHome = template.Must(template.ParseFiles("HTML/hpage.html"))      //templates2 -> tempHome            check      hpage -> homePage
var tempDetails = template.Must(template.ParseFiles("HTML/details.html")) //templates3 -> tempDetails    check      details -> details
var ApiElements []ArtistAPI                                               //ApiObject -> ApiElements    check

func artist(w http.ResponseWriter, r *http.Request) {

	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ApiData, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(ApiData, &ApiElements)

	DataArtists := ArtistsArray{
		Artists: ApiElements,
	}

	MapInt := map[string]interface{}{
		"DataArtists": DataArtists,
	}
	tempCard.Execute(w, MapInt)
}

func home(w http.ResponseWriter, r *http.Request) {

	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	var ApiElements ArtistAPI

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ApiDataArtist, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(ApiDataArtist, &ApiElements)

	tempHome.Execute(w, err)
}

//8
func details(w http.ResponseWriter, r *http.Request) {

	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	var LocationsObject Relation

	DataArtists := Descritpion{
		ID:           ApiElements[pathIDint-1].ID,
		Image:        ApiElements[pathIDint-1].Image,
		Members:      ApiElements[pathIDint-1].Members,
		CreationDate: ApiElements[pathIDint-1].CreationDate,
		FirstAlbum:   ApiElements[pathIDint-1].FirstAlbum,
		Locations:    ApiElements[pathIDint-1].Locations,
		ConcertDates: ApiElements[pathIDint-1].ConcertDates,
		Relations:    ApiElements[pathIDint-1].Relations,
	}

	Oui, err := http.Get(DataArtists.Relations)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	OuiData, err := ioutil.ReadAll(Oui.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(OuiData, &LocationsObject)

	MapInt := map[string]interface{}{
		"DataArtists": DataArtists,
		"Relation":    LocationsObject,
	}

	tempDetails.Execute(w, MapInt)
}

//9
func main() {
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/artist/", details)

	log.Fatal(http.ListenAndServe(":9090", nil))
}
