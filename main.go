package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {

	tmpl = template.Must(template.New("").Parse(page))

	http.HandleFunc("/say", cowServe)

	log.Println("Server started, listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
