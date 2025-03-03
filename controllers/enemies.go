package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/service"
	"main.go/utils"
)

type DataEnemies struct {
	Data    []service.Enemies
	IsLogin bool
}

type DataEnemieDetails struct {
	Data    service.Enemies
	IsLogin bool
}

func Enemies(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		Data := DataEnemies{
			Data:    AllDataOfAPI.AllEnnemies,
			IsLogin: IsLogin,
		}
		err1 := Templates.ExecuteTemplate(w, "enemies", Data)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}
func EnemiesDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			Id := r.URL.Query().Get("id")
			if !utils.HasIdInEnemie(Id, AllDataOfAPI.AllEnnemies) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				enemie, err := service.GetAllDetailsAboutOneEnemie(Id)
				if err != nil {
					log.Fatal(err)
				}
				Data := DataEnemieDetails{
					Data:    enemie,
					IsLogin: IsLogin,
				}
				err = Templates.ExecuteTemplate(w, "enemiesDetails", Data)
				if err != nil {
					http.Error(w, "Error rendering template", http.StatusInternalServerError)
				}
			}
		}
	}
}
