package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "homePage")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contactPage")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t , err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)
	http.ListenAndServe(":8080", nil)
	fmt.Print("The serveur start on port 8080 🔥")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

