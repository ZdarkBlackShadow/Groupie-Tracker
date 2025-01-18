package server

import (
	"log"
	"net/http"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		TimeWhenConnect = time.Now()
		IsLoad = true
		Progress = 0
		go LoadAllData()
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
