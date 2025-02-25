package controllers

import (
	"log"
	"net/http"
)

func AboutControllers(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "about", nil)
	if err != nil {
		log.Fatal(err)
	}
}
