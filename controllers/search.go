package controllers

import (
	"log"
	"net/http"
	"strconv"

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
	UserSearch, pageParam := r.URL.Query().Get(("query")), r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	ListOfSearchresult := utils.Search(UserSearch, AllDataOfAPI)
	pagedData, totalPages := utils.PaginateSearchResult(ListOfSearchresult, page, 20)
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
