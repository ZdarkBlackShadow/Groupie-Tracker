package server

import (
	"log"
	"net/http"

	"main.go/service"
)

func Domains(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Domain
	}
	Data := DataArtifacts{
		Data: service.GetAllDomainsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "domains", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func DomainDetail(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneDomains(Id)
	err := Templates.ExecuteTemplate(w, "domainDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
