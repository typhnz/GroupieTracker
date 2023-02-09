package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	Dates []string `json:"dates"`

}


type locations struct {
	ID int `json:"id"`
	Locations []string `json:"locations"`
}
