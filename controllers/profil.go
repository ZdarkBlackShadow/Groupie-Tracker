package controllers

import (
	"net/http"
	"time"
)

type ProfilStruct struct {
	IsLogin bool
}

func Profil(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !IsLogin {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		Data := ProfilStruct{
			IsLogin: IsLogin,
		}
		Templates.ExecuteTemplate(w, "profil", Data)
	}
}
