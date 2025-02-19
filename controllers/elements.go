package controllers

import (
	"log"
	"net/http"

	"main.go/service"
)

type DataElements struct {
	Data    []service.Element
	IsLogin bool
}

type DataElementDetails struct {
	Data    service.Element
	IsLogin bool
}

func Elements(w http.ResponseWriter, r *http.Request) {
	Data := DataElements{
		Data:    AllDataOfAPI.AllElement,
		IsLogin: IsLogin,
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
	element, err := service.GetAllDetailsAboutOneElements(Id)
	if err != nil {
		log.Fatal(err)
	}
	Data := DataElementDetails{
		Data:    element,
		IsLogin: IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "ElementDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
