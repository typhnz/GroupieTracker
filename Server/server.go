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

func mainPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "mainPage")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", http.FileServer(http.Dir("../templates/cssFile"))))
	http.Handle("/javaFile/", http.StripPrefix("/javaFile/", http.FileServer(http.Dir("../templates/javaFile"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("../templates/picture"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/mainPage", mainPage)
	http.ListenAndServe(":8080", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
