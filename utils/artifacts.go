package utils

import (
	"math"
	"net/http"
	"sort"
	"strconv"

	"main.go/service"
)

type ArtifactFilters struct {
	Rarity []string
	SortBy string
}

func GetArtifactFilters(r *http.Request) ArtifactFilters {
	filters := ArtifactFilters{
		Rarity: r.URL.Query()["rarity_filter"],
		SortBy: r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplyArtifactFilters(filters ArtifactFilters, AllArtifacts []service.ArtifactDetails) []service.ArtifactDetails {
	artifacts := []service.ArtifactDetails{}
	if len(filters.Rarity) == 1 {
		rarity, err := strconv.Atoi(filters.Rarity[0])
		if err != nil {
			return artifacts
		}
		artifacts = filterArtifactsByRarity(AllArtifacts, rarity)
	} else if len(filters.Rarity) == 2 {
		rarity, err := strconv.Atoi(filters.Rarity[0])
		if err != nil {
			return artifacts
		}
		artifacts = filterArtifactsByRarity(AllArtifacts, rarity)
		rarity, err = strconv.Atoi(filters.Rarity[1])
		if err != nil {
			return artifacts
		}
		artifacts = append(artifacts, filterArtifactsByRarity(AllArtifacts, rarity)...)
	} else {
		artifacts = AllArtifacts
	}
	if filters.SortBy != "" {
		switch filters.SortBy {
		case "az":
			artifacts = filterArtifactsByAZ(artifacts)
		case "za":
			artifacts = filterArtifactsByZA(artifacts)
		}
	}
	return artifacts
}

func filterArtifactsByAZ(artifacts []service.ArtifactDetails) []service.ArtifactDetails {
	sortedArtifacts := make([]service.ArtifactDetails, len(artifacts))
	copy(sortedArtifacts, artifacts)

	sort.Slice(sortedArtifacts, func(i, j int) bool {
		return sortedArtifacts[i].Name < sortedArtifacts[j].Name
	})
	return sortedArtifacts
}

func filterArtifactsByZA(artifacts []service.ArtifactDetails) []service.ArtifactDetails {
	sortedArtifacts := make([]service.ArtifactDetails, len(artifacts))
	copy(sortedArtifacts, artifacts)

	sort.Slice(sortedArtifacts, func(i, j int) bool {
		return sortedArtifacts[i].Name > sortedArtifacts[j].Name
	})
	return sortedArtifacts
}

func filterArtifactsByRarity(artifacts []service.ArtifactDetails, rarity int) []service.ArtifactDetails {
	filteredArtifacts := make([]service.ArtifactDetails, 0)
	for _, artifact := range artifacts {
		if artifact.MaxRarity == rarity {
			filteredArtifacts = append(filteredArtifacts, artifact)
		}
	}
	return filteredArtifacts
}

func ArtifactFilterChecker(filters ArtifactFilters) bool {
	if len(filters.Rarity) > 3 {
		return false
	}
	for _, v := range filters.Rarity {
		if !(v == "3" || v == "4" || v == "5" || v == "") {
			return false
		}
	}
	return (filters.SortBy == "az" || filters.SortBy == "za" || filters.SortBy == "")
}

func PaginateArtifacts(data []service.ArtifactDetails, page, itemsPerPage int) ([]service.ArtifactDetails, int, int) {
	totalItems := len(data)
	totalPages := int(math.Ceil(float64(totalItems) / float64(itemsPerPage)))
	if page > totalPages {
		page = totalPages
	}
	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage

	if start > totalItems {
		start = totalItems
	}
	if end > totalItems {
		end = totalItems
	}

	return data[start:end], totalPages, page
}
