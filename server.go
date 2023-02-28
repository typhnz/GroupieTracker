package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
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
	//Relation []string
}

type artistsData struct {
	Artists []Artist
}

const port = ":8080"

func main() {
	http.HandleFunc("/artists", displayArtists)
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates/cssFile"))))
	http.ListenAndServe(":8080", nil)
}

func artists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	var a []Artist
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func home(w http.ResponseWriter, r *http.Request) {
	data, err := artists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "home", data)
}

/*func artist(w http.ResponseWriter, r *http.Request) {
	data, err := artists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "artist", data)
}*/

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", nil)
}

func displayArtists(w http.ResponseWriter, r *http.Request) {
	data, err := artists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "artist", data)

	/*jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)*/
}
