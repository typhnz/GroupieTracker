package main

import (
	"encoding/json"
	"fmt"
)

type artist struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Members     int    `json:"members"`
	ReleaseDate string `json:"releasedate"`
}

func main() {
	u, err := json.Marshal(artist{Name: "Marina", ReleaseDate: "January 8th 2000"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
}
