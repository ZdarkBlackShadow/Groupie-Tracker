package controllers

import (
	"log"
	"net/http"

	"main.go/service"
)

type DataBoss struct {
	Data    []service.BossStruct
	IsLogin bool
}

type DataBossDetails struct {
	Data    service.BossStruct
	IsLogin bool
}

func Boss(w http.ResponseWriter, r *http.Request) {
	Data := DataBoss{
		Data:    AllDataOfAPI.AllBoss,
		IsLogin: IsLogin,
	}
	err1 := Templates.ExecuteTemplate(w, "boss", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func BossDetails(w http.ResponseWriter, r *http.Request) {
	var err error
	var boss service.BossStruct
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	boss, err = service.GetAllDetailsAboutOneBoss(Id)
	if err != nil {
		log.Fatal(err)
	}
	Data := DataBossDetails{
		Data:    boss,
		IsLogin: IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "bossDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
