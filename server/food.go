package server

import (
	"net/http"

	"main.go/service"
)

type FoodStructData struct {
	Data    map[string]service.FoodStruct
	IsLogin bool
}

func Food(w http.ResponseWriter, r *http.Request) {
	Data := FoodStructData{
		Data:    API_Data.AllFood,
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "food", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
