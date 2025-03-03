package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/service"
	"main.go/utils"
)

type PotionsStructData struct {
	Data    []service.Potion
	IsLogin bool
}

type PotionDetailStructData struct {
	Data    service.Potion
	IsLogin bool
}

func Potions(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		Data := PotionsStructData{
			Data:    AllDataOfAPI.AllPotions,
			IsLogin: IsLogin,
		}
		err := Templates.ExecuteTemplate(w, "potions", Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PotionsDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			Id := r.URL.Query().Get("id")
			if !utils.HasIdInPotions(Id, AllDataOfAPI.AllPotions) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				Data := PotionDetailStructData{
					Data:    service.GetDetailsAboutOnePotion(Id, AllDataOfAPI.AllPotions),
					IsLogin: IsLogin,
				}
				err := Templates.ExecuteTemplate(w, "potionsdetails", Data)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
