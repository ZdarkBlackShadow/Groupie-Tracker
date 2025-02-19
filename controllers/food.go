package controllers

import (
	"net/http"
	"strconv"

	"main.go/service"
	"main.go/utils"
)

type FoodStructData struct {
	Data        []service.FoodStruct
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type FoodDetailStructData struct {
	Data    service.FoodStruct
	IsLogin bool
}

func Food(w http.ResponseWriter, r *http.Request) {
	filters := utils.GetFoodFilters(r)
	DataFood := utils.ApplyFoodFilters(filters, AllDataOfAPI.AllFood)
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	ItemsPerPage := 16
	Data, totalPages := utils.PaginateFood(DataFood, page, ItemsPerPage)
	DataForTheTemplate := FoodStructData{
		Data:        Data,
		TotalPages:  totalPages,
		CurrentPage: page,
		IsLogin:     IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "food", DataForTheTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FoodDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Data := FoodDetailStructData{
		Data:    service.GetDetailsOfFood(id, AllDataOfAPI.AllFood),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "foodDetails", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
