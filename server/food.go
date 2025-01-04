package server

import (
	"net/http"

	"main.go/service"
)

func Food(w http.ResponseWriter, r *http.Request) {
	Data := service.GetAllDetailsOfFood()
	err := Templates.ExecuteTemplate(w, "food", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
