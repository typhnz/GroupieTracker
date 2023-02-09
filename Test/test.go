package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type artist struct {
	ID          int
	Image       string
	Name        string
	Members	 []string
	CreationDate int
	FirstAlbum string
	Locations string
	ConcertDates string
	Relation string
}

func main() {
	artists()
}

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
	fmt.Println("Personne:", a[0].Relation)
}
