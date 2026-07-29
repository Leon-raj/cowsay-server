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
	http.HandleFunc("/download", downloadHandler)

	log.Println("Server started, listening on http://localhost:8080/say")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
