package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
