package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"main.go/utils"
)

type DataSearchStruct struct {
	Data        []utils.SearchResultStruct
	Query       string
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

func Search(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if !r.URL.Query().Has("query") {
			ErrorCodeToSend.Update(400, "Bad request", "query is missing, please check the parameter of the URL of your request", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			UserSearch, pageParam, filters := r.URL.Query().Get(("query")), r.URL.Query().Get("page"), r.URL.Query()["sort_filter"]
			page, err := strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				page = 1
			}
			if !utils.FiltersSearchChecker(filters) {
				ErrorCodeToSend.Update(400, "Bad request", "Invalid parameter", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				ListOfSearchresult := utils.Search(UserSearch, AllDataOfAPI)
				ListToSend := utils.ApplySearchFilters(ListOfSearchresult, filters)
				pagedData, totalPages := utils.PaginateSearchResult(ListToSend, page, 20)
				Data := DataSearchStruct{
					Data:        pagedData,
					Query:       UserSearch,
					TotalPages:  totalPages,
					CurrentPage: page,
					IsLogin:     IsLogin,
				}
				err = Templates.ExecuteTemplate(w, "ResultOfTheSearch", Data)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
