package main

import (
	"bytes"
	"html/template"
	"net/http"

	"nmyk.io/cowsay"
)

func main() {

	tmpl := template.Must(template.New("").Parse(page))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		text := r.URL.Query().Get("text")
		if text == "" {
			text = "Linux is awesome!"
		}

		var b bytes.Buffer

		cowsay.Cow{}.Write(&b, []byte(text), false)

		tmpl.Execute(w, map[string]any{
			"Input": text,
			"Cow":   b.String(),
		})
	})

	http.ListenAndServe(":8080", nil)
}
