package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/service"
	"main.go/utils"
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
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		Data := DataElements{
			Data:    AllDataOfAPI.AllElement,
			IsLogin: IsLogin,
		}
		err1 := Templates.ExecuteTemplate(w, "elements", Data)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func ElementDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			Id := r.URL.Query().Get("id")
			if !utils.HasIsInElement(Id, AllDataOfAPI.AllElement) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
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
		}
	}
}
