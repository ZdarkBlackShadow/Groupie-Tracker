package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

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
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		filters := utils.GetWeaponFilters(r)
		if !utils.WeaponsFiltersChecker(filters) {
			ErrorCodeToSend.Update(400, "Bad request", "Invalid filter", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
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
}

func WeaponDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			Id := r.URL.Query().Get("id")
			if !utils.HasIdInWeapons(Id, AllDataOfAPI.AllWeapons) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				Data := DataWeaponDetails{
					Data:    service.GetAllDataAboutOneWeapon(Id),
					IsLogin: IsLogin,
				}
				err := Templates.ExecuteTemplate(w, "weaponsDetails", Data)
				if err != nil {
					http.Error(w, "Error rendering template", http.StatusInternalServerError)
				}
			}
		}
	}
}
