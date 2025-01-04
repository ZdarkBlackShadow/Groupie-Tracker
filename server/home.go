package server

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err1 := Templates.ExecuteTemplate(w, "home", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}
