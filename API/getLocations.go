package location

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func location() {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}