package server

import (
	"log"
	"net/http"

	"main.go/service"
)

func Elements(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Element
	}
	Data := DataArtifacts{
		Data: service.GetAllElementsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "elements", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func ElementDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneElements(Id)
	err := Templates.ExecuteTemplate(w, "ElementDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
