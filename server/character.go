package server

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"main.go/service"
)

type DataCharacters struct {
	Data        []service.Characters
	TotalPages  int
	CurrentPage int
}

func PaginateCharacters(data []service.Characters, page, itemsPerPage int) ([]service.Characters, int) {
	totalItems := len(data)
	totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))

	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage

	if start > totalItems {
		start = totalItems
	}
	if end > totalItems {
		end = totalItems
	}

	return data[start:end], totalPages
}
func Characters(w http.ResponseWriter, r *http.Request) {
	allCharacters := service.GetAllCharactersDetails()

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	itemsPerPage := 20
	pagedData, totalPages := PaginateCharacters(allCharacters, page, itemsPerPage)

	dataArtifacts := DataCharacters{
		Data:        pagedData,
		TotalPages:  totalPages,
		CurrentPage: page,
	}
	Templates.ExecuteTemplate(w, "characters", dataArtifacts)
}

func CharactersDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneCharacters(Id)
	err := Templates.ExecuteTemplate(w, "characterDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}
