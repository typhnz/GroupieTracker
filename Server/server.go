package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("The serveur start on port 8080 🔥")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenue sur mon serveur Go!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

