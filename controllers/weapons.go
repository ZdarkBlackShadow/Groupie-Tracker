package controllers

import (
	"log"
	"net/http"
	"strconv"

	"main.go/service"
	"main.go/utils"
)

type DataWeapons struct {
	Data        []service.Weapon
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type DataWeaponDetails struct {
	Data    service.Weapon
	IsLogin bool
}

func Weapons(w http.ResponseWriter, r *http.Request) {
	filters := utils.GetWeaponFilters(r)
	if !utils.WeaponsFiltersChecker(filters) {
		http.Redirect(w, r, "/400", http.StatusSeeOther)
	} else {
		AllWeopns := utils.ApplyWeaponFilters(filters, AllDataOfAPI.AllWeapons)

		pageParam := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			page = 1
		}
		pagedData, totalPages, page := utils.PaginateWeapons(AllWeopns, page, 20)
		Data := DataWeapons{
			Data:        pagedData,
			TotalPages:  totalPages,
			CurrentPage: page,
			IsLogin:     IsLogin,
		}
		err1 := Templates.ExecuteTemplate(w, "weapon", Data)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}

func WeaponDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := DataWeaponDetails{
		Data:    service.GetAllDataAboutOneWeapon(Id),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "weaponsDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
