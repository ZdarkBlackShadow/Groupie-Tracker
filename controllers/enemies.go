package controllers

import (
	"log"
	"net/http"

	"main.go/service"
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
	Data := DataEnemies{
		Data:    AllDataOfAPI.AllEnnemies,
		IsLogin: IsLogin,
	}
	err1 := Templates.ExecuteTemplate(w, "enemies", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}
func EnemiesDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
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
