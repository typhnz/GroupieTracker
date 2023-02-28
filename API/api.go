package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relation     string   `json:"relations"`
}

type artistarray struct {
	Array []artist
}
const port = ":8080"

var artistData artistarray

func artists() {
	var a []artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &artistData.Array)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(a); i++ {
		fmt.Println("Image:", a[i].Image)
		fmt.Println("Nom:", a[i].Name)
		fmt.Println("Membres:", a[i].Members)
		fmt.Println("Date of creation:", a[i].CreationDate)
		fmt.Println("Date of first Album:", a[i].FirstAlbum)
		fmt.Println("Location:", a[i].Locations)
		fmt.Println("Date of concert:", a[i].ConcertDates)
		fmt.Println("\n")
	}
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/mainPage.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, artistData)
}
func main() {
	artists()
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", http.FileServer(http.Dir("../templates/cssFile"))))
	http.Handle("/javaFile/", http.StripPrefix("/javaFile/", http.FileServer(http.Dir("../templates/javaFile"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("../templates/picture"))))
	// http.HandleFunc("/", home)
	// http.HandleFunc("/contact", contact)
	http.HandleFunc("/mainPage", mainPage)
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
