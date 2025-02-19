package controllers

import (
	"fmt"
	"net/http"

	"main.go/utils"
)

func Loading(w http.ResponseWriter, r *http.Request) {

	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	if !IsLoad {
		IsLoad = true
		utils.Progress = 0
		go AllDataOfAPI.LoadAllData()
		http.Redirect(w, r, "/loading", http.StatusSeeOther)
		return
	}
	err := Templates.ExecuteTemplate(w, "data", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProgressF(w http.ResponseWriter, r *http.Request) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"progress": %d}`, utils.Progress)))
}
