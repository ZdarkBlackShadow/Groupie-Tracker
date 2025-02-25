package controllers

import (
	"log"
	"net/http"

	"main.go/service"
	"main.go/utils"
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
	data := DataDomains{
		Data:    AllDataOfAPI.AllDomain,
		IsLogin: IsLogin,
	}
	err1 := Templates.ExecuteTemplate(w, "domains", data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func DomainDetail(w http.ResponseWriter, r *http.Request) {
	var (
		id     string
		domain service.Domain
		err    error
		data   DataDomainDetails
	)
	if !r.URL.Query().Has("id") {
		ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	} else {
		id = r.URL.Query().Get("id")
		if !utils.HasIdInDomain(id, AllDataOfAPI.AllDomain) {
			ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			domain, err = service.GetAllDetailsAboutOneDomains(id)
			if err != nil {
				log.Fatal(err)
			}
			data = DataDomainDetails{
				Data:    domain,
				IsLogin: IsLogin,
			}
			err = Templates.ExecuteTemplate(w, "domainDetails", data)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
		}
	}
}
