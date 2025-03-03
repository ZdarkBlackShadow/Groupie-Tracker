package controllers

import (
	"log"
	"net/http"
	"time"
)

func AboutControllers(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	err := Templates.ExecuteTemplate(w, "about", nil)
	if err != nil {
		log.Fatal(err)
	}
}
