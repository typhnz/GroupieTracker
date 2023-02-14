package GroupieTracker

import (
	"encoding/json"
	"fmt"
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

func ArtistFunction(element Artist) Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal((body), &element)
	if err != nil {
		fmt.Println("Error:", err)
	}
	/*for i := 0; i < len(element); i++ {
		fmt.Println("Image:", element[i].Image)
		fmt.Println("Name:", element[i].Name)
		fmt.Println("Members:", element[i].Members)
		fmt.Println("CreationDate:", element[i].CreationDate)
		fmt.Println("FirstAlbum", element[i].FirstAlbum)
		fmt.Println("Locations:", element[i].Locations)
		fmt.Println("ConcertDates:", element[i].ConcertDates)
		fmt.Printf("\n")
	}*/
	return (element)
}
