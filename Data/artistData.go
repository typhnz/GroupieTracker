package data

type artist struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Members     int    `json:"members"`
	ReleaseDate string `json:"releasedate"`
}
