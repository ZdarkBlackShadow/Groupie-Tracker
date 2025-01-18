package server

import "net/http"

func Profil(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "profil", nil)
}
