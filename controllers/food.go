package controllers

import (
	"net/http"
	"strconv"
	"time"

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
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		filters := utils.GetFoodFilters(r)
		if !utils.CheckFoodFilter(filters) {
			ErrorCodeToSend.Update(400, "Bad Request", "Error in the parameter of the URL", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
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
}

func FoodDetail(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			id := r.URL.Query().Get("id")
			if !utils.HasIdInFood(id, AllDataOfAPI.AllFood) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			}
			Data := FoodDetailStructData{
				Data:    service.GetDetailsOfFood(id, AllDataOfAPI.AllFood),
				IsLogin: IsLogin,
			}
			err := Templates.ExecuteTemplate(w, "foodDetails", Data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
