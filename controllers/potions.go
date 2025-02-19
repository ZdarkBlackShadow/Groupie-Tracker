package controllers

import (
	"log"
	"net/http"

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
	Data := PotionsStructData{
		Data:    AllDataOfAPI.AllPotions,
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "potions", Data)
	if err != nil {
		log.Fatal(err)
	}
}

func PotionsDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if utils.HasIdInPotions(Id, AllDataOfAPI.AllPotions) {
		Data := PotionDetailStructData{
			Data:    service.GetDetailsAboutOnePotion(Id, AllDataOfAPI.AllPotions),
			IsLogin: IsLogin,
		}
		err := Templates.ExecuteTemplate(w, "potionsdetails", Data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//gestion d'erreur
	}
}
