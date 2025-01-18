package server

import "net/http"

func Logout(w http.ResponseWriter, r *http.Request) {
	IsLogin = false
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
