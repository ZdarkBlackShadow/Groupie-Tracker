package server

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"main.go/service"
)

type DataArtifacts struct {
	Data        []service.ArtifactDetails
	TotalPages  int
	CurrentPage int
}

func PaginateArtifacts(data []service.ArtifactDetails, page, itemsPerPage int) ([]service.ArtifactDetails, int) {
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

func Artifacts(w http.ResponseWriter, r *http.Request) {
	allArtifacts := service.GetAllArtifactsDetails()

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	pagedData, totalPages := PaginateArtifacts(allArtifacts, page, 15)
	Data := DataArtifacts{
		Data:        pagedData,
		TotalPages:  totalPages,
		CurrentPage: page,
	}
	err1 := Templates.ExecuteTemplate(w, "artifacts", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func ArtifactsDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDataAboutOneArtifact(Id)
	err := Templates.ExecuteTemplate(w, "ArtifactDetail", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
