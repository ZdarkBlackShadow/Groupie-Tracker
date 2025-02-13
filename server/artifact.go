package server

import (
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"

	"main.go/service"
)

type DataArtifacts struct {
	Data        []service.ArtifactDetails
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type DataArtifactDetails struct {
	Data    service.ArtifactDetails
	IsLogin bool
}

type ArtifactFilters struct {
	Rarity []string
	SortBy string
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

func GetArtifactFilters(r *http.Request) ArtifactFilters {
	filters := ArtifactFilters{
		Rarity: r.URL.Query()["rarity_filter"],
		SortBy: r.URL.Query().Get("sort_by"),
	}
	return filters
}

func ApplyArtifactFilters(filters ArtifactFilters) []service.ArtifactDetails {
	artifacts := []service.ArtifactDetails{}
	if len(filters.Rarity) == 1 {
		rarity, err := strconv.Atoi(filters.Rarity[0])
		if err != nil {
			return artifacts
		}
		artifacts = FilterArtifactsByRarity(API_Data.ALLArtifacts, rarity)
	} else if len(filters.Rarity) == 2 {
		rarity, err := strconv.Atoi(filters.Rarity[0])
		if err != nil {
			return artifacts
		}
		artifacts = FilterArtifactsByRarity(API_Data.ALLArtifacts, rarity)
		rarity, err = strconv.Atoi(filters.Rarity[1])
		if err != nil {
			return artifacts
		}
		artifacts = append(artifacts, FilterArtifactsByRarity(API_Data.ALLArtifacts, rarity)...)
	} else {
		artifacts = API_Data.ALLArtifacts
	}
	if filters.SortBy == "" {
		switch filters.SortBy {
		case "az":
			artifacts = FilterArtifactsByAZ(artifacts)
		case "za":
			artifacts = FilterArtifactsByZA(artifacts)
		}
	}
	return artifacts
}

func Artifacts(w http.ResponseWriter, r *http.Request) {
	allArtifacts := ApplyArtifactFilters(GetArtifactFilters(r))

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	pagedData, totalPages := PaginateArtifacts(allArtifacts, page, 16)
	Data := DataArtifacts{
		Data:        pagedData,
		TotalPages:  totalPages,
		CurrentPage: page,
		IsLogin:     IsLogin,
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
	Data := DataArtifactDetails{
		Data:    service.GetAllDataAboutOneArtifact(Id),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "ArtifactDetail", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func FilterArtifactsByAZ(artifacts []service.ArtifactDetails) []service.ArtifactDetails {
	sortedArtifacts := make([]service.ArtifactDetails, len(artifacts))
	copy(sortedArtifacts, artifacts)

	sort.Slice(sortedArtifacts, func(i, j int) bool {
		return sortedArtifacts[i].Name < sortedArtifacts[j].Name
	})
	return sortedArtifacts
}

func FilterArtifactsByZA(artifacts []service.ArtifactDetails) []service.ArtifactDetails {
	sortedArtifacts := make([]service.ArtifactDetails, len(artifacts))
	copy(sortedArtifacts, artifacts)

	sort.Slice(sortedArtifacts, func(i, j int) bool {
		return sortedArtifacts[i].Name > sortedArtifacts[j].Name
	})
	return sortedArtifacts
}

func FilterArtifactsByRarity(artifacts []service.ArtifactDetails, rarity int) []service.ArtifactDetails {
	filteredArtifacts := make([]service.ArtifactDetails, 0)
	for _, artifact := range artifacts {
		if artifact.MaxRarity == rarity {
			filteredArtifacts = append(filteredArtifacts, artifact)
		}
	}
	return filteredArtifacts
}
