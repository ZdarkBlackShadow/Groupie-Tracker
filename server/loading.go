package server

import "net/http"

func Loading(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "data", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
