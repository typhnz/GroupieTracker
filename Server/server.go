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

type ExtractLocation struct {
	Index []Locations `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type ExtractDate struct {
	Index []Dates `json:"index"`
}

type ArtistsArray struct {
	Artists   []ArtistAPI
	Relation  ExtractRelation
	Locations ExtractLocation
	Dates     ExtractDate
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

var artistsData ArtistsArray

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

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func Arstists(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "artist", artistsData)
}

func details(w http.ResponseWriter, r *http.Request) {
	click := r.FormValue("true")
	id, _ := strconv.Atoi(click)
	artistsData.Artists[id-1].Relations = artistsData.Relation.Index[id-1]
	renderTemplate(w, "details", artistsData.Artists[id-1])
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query != "" {
		fmt.Fprintf(w, "Vous avez cherch??: %s", query)
	} else {
		renderTemplate(w, "mainPage", nil)
	}
}

const port = ":8080"

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	GetAPI("artists")
	GetAPI("relation")
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", http.FileServer(http.Dir("../templates/cssFile"))))
	http.Handle("/javaFile/", http.StripPrefix("/javaFile/", http.FileServer(http.Dir("../templates/javaFile"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("../templates/picture"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/details", details)
	http.HandleFunc("/artist", Arstists)
	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
