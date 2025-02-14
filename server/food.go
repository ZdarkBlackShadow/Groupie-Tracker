package server

import (
	"net/http"

	"main.go/service"
)

type FoodStructData struct {
	Data    []service.FoodStruct
	IsLogin bool
}

type FoodDetailStructData struct {
	Data service.FoodStruct
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

func FoodDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Data := FoodDetailStructData{
		Data:    service.GetDetailsOfFood(id, API_Data.AllFood),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "foodDetails", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}