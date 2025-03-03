package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/service"
	"main.go/utils"
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
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		Data := DataBoss{
			Data:    AllDataOfAPI.AllBoss,
			IsLogin: IsLogin,
		}
		err1 := Templates.ExecuteTemplate(w, "boss", Data)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func BossDetails(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		boss service.BossStruct
		id   string
		data DataBossDetails
	)
	if !r.URL.Query().Has("id") {
		ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	} else {
		id = r.URL.Query().Get("id")
		if !utils.HasIdInBoss(id, AllDataOfAPI.AllBoss) {
			ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			boss, err = service.GetAllDetailsAboutOneBoss(id)
			if err != nil {
				log.Fatal(err)
			}
			data = DataBossDetails{
				Data:    boss,
				IsLogin: IsLogin,
			}
			err = Templates.ExecuteTemplate(w, "bossDetails", data)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
		}
	}
}
