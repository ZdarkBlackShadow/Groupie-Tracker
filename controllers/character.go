package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"main.go/service"
	"main.go/utils"
)

type DataCharacters struct {
	Data        []service.Characters
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type DataCharactersDetails struct {
	Data    service.Characters
	IsLogin bool
}

func Characters(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		filters := utils.GetCharacterFilters(r)
		if !utils.CharacterFiltersChecker(filters) {
			ErrorCodeToSend.Update(400, "Bad Request", "Error in the parameter of the URL", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			allCharacters := utils.ApplyCharacterFilters(filters, AllDataOfAPI.AllCharacters)
			pageParam := r.URL.Query().Get("page")
			page, err := strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				page = 1
			}

			itemsPerPage := 16
			pagedData, totalPages, page := utils.PaginateCharacters(allCharacters, page, itemsPerPage)

			dataArtifacts := DataCharacters{
				Data:        pagedData,
				TotalPages:  totalPages,
				CurrentPage: page,
				IsLogin:     IsLogin,
			}
			err = Templates.ExecuteTemplate(w, "characters", dataArtifacts)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				fmt.Println(err)
			}
		}
	}
}

func CharactersDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad request", "The id is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			Id := r.URL.Query().Get("id")
			if !utils.HasIdInCharacter(Id, AllDataOfAPI.AllCharacters) {
				ErrorCodeToSend.Update(400, "Bad request", "The id is invalid, please check the parameter of the URL of your request", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				var err error
				if Id == "" {
					http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
					return
				}
				character, err := service.GetAllDetailsAboutOneCharacters(Id)
				if err != nil {
					log.Fatal(err)
				}
				Datat := DataCharactersDetails{
					Data:    character,
					IsLogin: IsLogin,
				}
				err = Templates.ExecuteTemplate(w, "characterDetails", Datat)
				if err != nil {
					http.Error(w, "Error rendering template", http.StatusInternalServerError)
					fmt.Println(err)
				}
			}
		}
	}
}
