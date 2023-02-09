package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type artist struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Members     int    `json:"members"`
	ReleaseDate string `json:"releasedate"`
}

type dates struct {
	ID int `json:"id"`
	Date []string `json:"dates"`
}


type location struct {
	ID int `json:"id"`
	Locations []string `json:"locations"`
}

type relations struct {
	ID int `json:"id"`
	Relations []string `json:"relations"`
}

func main() {
	artists()
}

func artists() {
	var a artist
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
	fmt.Println("Personne:", a)
}

/*func dat() {
	var d dates
	url := "https://groupietrackers.herokuapp.com/api/dates"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Dates:", d)
}

func locations() {
	var l locations
	url := "https://groupietrackers.herokuapp.com/api/locations"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &l)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Locations:", l)
}

func relations() {
	var r relations
	url := "https://groupietrackers.herokuapp.com/api/relations"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &r)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Relations:", r)
}

func main() {
	artists()
	dates()
	locations()
}

*/