package controllers

import "net/http"

type ProfilStruct struct {
	IsLogin bool
}

func Profil(w http.ResponseWriter, r *http.Request) {
	if !IsLogin {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	Data := ProfilStruct{
		IsLogin: IsLogin,
	}
	Templates.ExecuteTemplate(w, "profil", Data)
}
