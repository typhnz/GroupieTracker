package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Relations    Relation
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"DatesLocations"`
}

type ExtractRelation struct {
	Index []Relation `json:"index"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     Dates
}

type Location struct {
	Index []Locations `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type blabla struct {
	Index []Dates `json:"index"`
}

type ArtistsArray struct {
	Artists   []ArtistAPI
	Relation  ExtractRelation
	Locations Location
	Dates     blabla
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

const port = ":8080"

var apiElements []ArtistAPI

var artistsData ArtistsArray

func details(w http.ResponseWriter, r *http.Request) {

	az := r.FormValue("Oui")
	fmt.Println(az)
	id, _ := strconv.Atoi(az)
	artistsData.Artists[id-1].Relations = artistsData.Relation.Index[id-1]
	renderTemplate(w, "details", artistsData.Artists[id-1])
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func GetAPI(pathAPI string) {
	restAPI := "https://groupietrackers.herokuapp.com/api/"

	response, err := http.Get(restAPI + pathAPI)

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	switch pathAPI {
	case "artists":
		json.Unmarshal(responseData, &artistsData.Artists)
	case "relation":
		json.Unmarshal(responseData, &artistsData.Relation)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func Arstists(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/artist.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, artistsData)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query != "" {
		fmt.Fprintf(w, "Vous avez cherché: %s", query)
	} else {
		renderTemplate(w, "mainPage", nil)
	}
}

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", http.FileServer(http.Dir("../templates/cssFile"))))
	http.Handle("/javaFile/", http.StripPrefix("/javaFile/", http.FileServer(http.Dir("../templates/javaFile"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("../templates/picture"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/details", details)
	http.HandleFunc("/artist", Arstists)
	http.HandleFunc("/search", searchHandler)
	GetAPI("artists")
	GetAPI("relation")
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
