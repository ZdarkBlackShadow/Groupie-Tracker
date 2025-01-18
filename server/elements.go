package server

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
		Data:    API_Data.AllElement,
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
	Data := DataElementDetails{
		Data:    service.GetAllDetailsAboutOneElements(Id),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "ElementDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
