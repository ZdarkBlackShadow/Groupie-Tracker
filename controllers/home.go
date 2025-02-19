package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		TimeWhenConnect = time.Now()
		IsLoad = true
		go AllDataOfAPI.LoadAllData()
		http.Redirect(w, r, "/loading", http.StatusSeeOther)
		return
	}
	type HomeData struct {
		IsLogin bool
	}
	Data := HomeData{IsLogin: IsLogin}

	err1 := Templates.ExecuteTemplate(w, "home", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}
