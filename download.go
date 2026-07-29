package main

import (
	"log"
	"net/http"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for download")
	http.Redirect(w, r, downloadURL, http.StatusMovedPermanently)
	log.Println("Redirecting to", downloadURL)
}
