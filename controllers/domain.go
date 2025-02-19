package controllers

import (
	"log"
	"net/http"

	"main.go/service"
)

type DataDomains struct {
	Data    []service.Domain
	IsLogin bool
}

type DataDomainDetails struct {
	Data    service.Domain
	IsLogin bool
}

func Domains(w http.ResponseWriter, r *http.Request) {
	Data := DataDomains{
		Data:    AllDataOfAPI.AllDomain,
		IsLogin: IsLogin,
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
	domain, err := service.GetAllDetailsAboutOneDomains(Id)
	if err != nil {
		log.Fatal(err)
	}
	Data := DataDomainDetails{
		Data:    domain,
		IsLogin: IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "domainDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
