package server

import (
	"fmt"
	"net/http"
)

func Loading(w http.ResponseWriter, r *http.Request) {

	Mu.Lock()
	defer Mu.Unlock()

	if !IsLoad {
		IsLoad = true
		Progress = 0
		go LoadAllData()
		http.Redirect(w, r, "/loading", http.StatusSeeOther)
		return
	}
	err := Templates.ExecuteTemplate(w, "data", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProgressF(w http.ResponseWriter, r *http.Request) {
	Mu.Lock()
	defer Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"progress": %d}`, Progress)))
}
