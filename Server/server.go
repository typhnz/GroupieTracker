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
<<<<<<< HEAD
	t, err := template.ParseFiles("templates/" + tmpl + ".page.tmpl")
=======
	t, err := template.ParseFiles("../templates/" + tmpl + ".html")
>>>>>>> c705692758414836dff95ee2ffb94cc7a5ad7a09
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, artists)
}

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", http.FileServer(http.Dir("../templates/cssFile"))))
	http.Handle("/javaFile/", http.StripPrefix("/javaFile/", http.FileServer(http.Dir("../templates/javaFile"))))
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
