package main

import (
	"fmt"
	"groupie/GroupieTracker"
	"html/template"
	"net/http"
)

const port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {
	type test struct {
		Name string
	}
	var Test test //creating an instance
	Test.Name = "rémi" //defining a variable in this instance to use for the template
	tmpl := template.Must(template.ParseFiles("./Templates/index.html")) //fetching the template html

	tmpl.Execute(w, Test) //executing template with the data

	var element g.Artist
	g.ArtistFunction(element)
	fmt.Println(element)

	//renderTemplate(w, g.artists(&element))
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	/*t, err := template.ParseFiles("../templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, element)*/
}

func main() {
	fmt.Println("(http://localhost:8080) - The serveur start on port", port)
	templates := http.FileServer(http.Dir("../templates/cssFile"))
	http.Handle("/cssFile/", http.StripPrefix("/cssFile/", templates))
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
