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
	"strings"
	"text/template"
)

// artists -> main
// hpage -> home
//

//api
type API struct {
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

//api
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"DatesLocations"`
}

//api
type Artists1 struct {
	Artists []API
}

//api
type DescritpionPage struct {
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

//api
type SearchBar struct {
	Artist    API
	SearchBar bool
}

//api
var templates = template.Must(template.ParseFiles("HTML/artists.html")) //1-> home 2->main 3->idCard
var templates2 = template.Must(template.ParseFiles("HTML/hpage.html"))
var templates3 = template.Must(template.ParseFiles("HTML/details.html"))
var ApiObject []API

//var data string
var Id_data string

//server
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

	json.Unmarshal(ApiData, &ApiObject)

	searchBar := r.FormValue("SearchBar")
	pouet := r.FormValue("filter")
	var SearchBar3 SearchBar
	fmt.Println(pouet)

	for i := 0; i < len(ApiObject); i++ {
		name := strings.ToUpper(ApiObject[i].Name)
		album := strings.ToUpper(ApiObject[i].FirstAlbum)
		creationDate := ApiObject[i].CreationDate
		creationDate2 := strconv.Itoa(creationDate)
		for z := 0; z < len(ApiObject[i].Members); z++ {
			members := strings.ToUpper(ApiObject[i].Members[z])

			searchBar = strings.ToUpper(searchBar)
			if name == searchBar || album == searchBar || creationDate2 == searchBar || members == searchBar {
				SearchBar3 = SearchBar{
					Artist:    ApiObject[i],
					SearchBar: true,
				}
			}
		}
	}

	VarArtists := Artists1{
		Artists: ApiObject,
	}

	MapInt := map[string]interface{}{
		"VarArtists": VarArtists,
		"SearchBar2": SearchBar3,
	}
	templates.Execute(w, MapInt)
}

//server
func home(w http.ResponseWriter, r *http.Request) {

	Api, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	var ApiObject API

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	ApiDataArtist, err := ioutil.ReadAll(Api.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(ApiDataArtist, &ApiObject)

	templates2.Execute(w, err)
}

//server
func details(w http.ResponseWriter, r *http.Request) {

	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	var LocationsObject Relation

	VarArtists := DescritpionPage{
		ID:           ApiObject[pathIDint-1].ID,
		Image:        ApiObject[pathIDint-1].Image,
		Members:      ApiObject[pathIDint-1].Members,
		CreationDate: ApiObject[pathIDint-1].CreationDate,
		FirstAlbum:   ApiObject[pathIDint-1].FirstAlbum,
		Locations:    ApiObject[pathIDint-1].Locations,
		ConcertDates: ApiObject[pathIDint-1].ConcertDates,
		Relations:    ApiObject[pathIDint-1].Relations,
	}

	Oui, err := http.Get(VarArtists.Relations)

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
		"VarArtists": VarArtists,
		"Relation":   LocationsObject,
	}

	templates3.Execute(w, MapInt)
}

//server
// func main() {
// 	fs := http.FileServer(http.Dir("/css"))
// 	http.Handle("/css/", http.StripPrefix("/css/", fs))
// 	http.HandleFunc("/", home)
// 	http.HandleFunc("/artist", artist)
// 	http.HandleFunc("/artist/", details)
// 	//on ajoute ici

// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }
