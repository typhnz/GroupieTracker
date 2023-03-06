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

const port = ":8080"

var apiElements []ArtistAPI

var artistsData ArtistsArray

func details(w http.ResponseWriter, r *http.Request) {
	az := r.FormValue("Oui")
	fmt.Println(az)
	id, _ := strconv.Atoi(az)
	fmt.Println(artistsData.Artists[id-1])
	// pathID := r.URL.Path
	// pathID = path.Base(pathID)
	// pathIDint, _ := strconv.Atoi(pathID)
	// var locationsObject Relation

	// dataArtists := Description{
	// 	ID:           apiElements[pathIDint-1].ID,
	// 	Image:        apiElements[pathIDint-1].Image,
	// 	Members:      apiElements[pathIDint-1].Members,
	// 	CreationDate: apiElements[pathIDint-1].CreationDate,
	// 	FirstAlbum:   apiElements[pathIDint-1].FirstAlbum,
	// 	Locations:    apiElements[pathIDint-1].Locations,
	// 	ConcertDates: apiElements[pathIDint-1].ConcertDates,
	// 	Relations:    apiElements[pathIDint-1].Relations,
	// }

	// relations, err := http.Get(dataArtists.Relations)

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// relationsData, err := ioutil.ReadAll(relations.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// json.Unmarshal(relationsData, &locationsObject)

	// // mapInt := map[string]interface{}{
	// // 	"DataArtists": dataArtists,
	// // 	"Relations":   locationsObject,
	// // }
	
	//var locationsObject Relation


	fmt.Println(artistsData.Artists[id-1].Relations)
	//fmt.Println(locationsObject)
	renderTemplate(w, "details", artistsData.Artists[id-1])
	renderTemplate(w, "details", artistsData.Artists[id-1].Relations)

	//json.Unmarshal([]byte(artistsData.Artists[id-1].Relations), &locationsObject)
}

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

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

	artistsData.Artists = apiElements

	renderTemplate(w, "artist", artistsData)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
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
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
