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

func api(w http.ResponseWriter, R *http.Request) {
	renderTemplate(w, "api")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, artists)
}

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	templates := http.FileServer(http.Dir("../templates/cssFile"))
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", templates))
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/api", api)
	http.HandleFunc("/renderTemplate", renderTemplate)
	http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
