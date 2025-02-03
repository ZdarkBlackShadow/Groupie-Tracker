package server

import (
	"log"
	"net/http"

	"main.go/service"
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
		Data:    service.GetAllPotions(),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "potions", Data)
	if err != nil {
		log.Fatal(err)
	}
}

func PotionsDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := PotionDetailStructData{
		Data:    service.GetDetailsAboutOnePotion(Id, API_Data.AllPotions),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "potionsdetails", Data)
	if err != nil {
		log.Fatal(err)
	}
}
