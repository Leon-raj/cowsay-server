package main

import (
	"bytes"
	"log"
	"net/http"

	"nmyk.io/cowsay"
)

func cowServe(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	log.Printf("Received GET request for text = %s\n", text)
	if text == "" {
		text = "Linux is awesome!"
	}

	var b bytes.Buffer
	cowsay.Cow{}.Write(&b, []byte(text), false)
	tmpl.Execute(w, map[string]any{
		"Input": text,
		"Cow":   b.String(),
	})

	log.Println("Handled request")
}
