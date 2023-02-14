package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	//Relation []string
}

var element []artist

func artists() {
	var a []artist
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &a)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
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