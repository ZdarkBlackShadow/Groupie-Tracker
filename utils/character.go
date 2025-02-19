package utils

import (
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"main.go/service"
)

type CharacterFilters struct {
	Gender            []string
	Nations           []string
	Rarity            []string
	Element           []string
	ReleaseDateFilter string
}

func PaginateCharacters(data []service.Characters, page, itemsPerPage int) ([]service.Characters, int, int) {
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

func GetCharacterFilters(r *http.Request) CharacterFilters {
	filters := CharacterFilters{
		Gender:            r.URL.Query()["gender_filter"],
		Nations:           r.URL.Query()["nations_filter"],
		Rarity:            r.URL.Query()["rarity_filter"],
		Element:           r.URL.Query()["elements_filter"],
		ReleaseDateFilter: r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplyCharacterFilters(filters CharacterFilters, AllCharacters []service.Characters) []service.Characters {
	Characters := []service.Characters{}
	if len(filters.Gender) == 1 {
		if filters.Gender[0] == "male" {
			MaleCharacters := filterCharactersByMale(AllCharacters)
			Characters = append(Characters, MaleCharacters...)
		} else {
			FemaleCharacters := filterCharactersByFemale(AllCharacters)
			Characters = append(Characters, FemaleCharacters...)
		}
	}
	if len(filters.Nations) > 0 {
		NationCharacters := filterCharactersByNation(filters.Nations, AllCharacters)
		Characters = append(Characters, NationCharacters...)
	}
	if len(filters.Rarity) == 1 {
		for _, rarity := range filters.Rarity {
			rarityInt, err := strconv.Atoi(rarity)
			if err != nil {
				continue
			}
			RarityCharacters := filterCharactersByRarity(rarityInt, AllCharacters)
			Characters = append(Characters, RarityCharacters...)
		}
	}
	if len(filters.Element) > 0 {
		VisionCharacters := filterCharactersByVision(filters.Element, AllCharacters)
		Characters = append(Characters, VisionCharacters...)
	}
	if len(Characters) == 0 {
		Characters = AllCharacters
	} else {
		Characters = removeDuplicatesOfCharacters(Characters)
	}
	if filters.ReleaseDateFilter != "" {
		if filters.ReleaseDateFilter == "recent" {
			Characters = sortCharactersRecent(Characters)
		} else if filters.ReleaseDateFilter == "oldest" {
			Characters = sortCharactersByOldest(Characters)
		} else if filters.ReleaseDateFilter == "az" {
			Characters = filterCharactersByAZ(Characters)
		} else if filters.ReleaseDateFilter == "za" {
			Characters = filterCharactersByZA(Characters)
		}
	}
	return Characters
}

func removeDuplicatesOfCharacters(characters []service.Characters) []service.Characters {
	encountered := map[string]bool{}
	result := []service.Characters{}

	for _, character := range characters {
		if !encountered[character.Id] {
			encountered[character.Id] = true
			result = append(result, character)
		}
	}
	return result
}

func filterCharactersByAZ(characters []service.Characters) []service.Characters {
	sortedCharacters := make([]service.Characters, len(characters))
	copy(sortedCharacters, characters)

	sort.Slice(sortedCharacters, func(i, j int) bool {
		return sortedCharacters[i].Name < sortedCharacters[j].Name
	})
	return sortedCharacters
}

func filterCharactersByZA(characters []service.Characters) []service.Characters {
	sortedCharacters := make([]service.Characters, len(characters))
	copy(sortedCharacters, characters)

	sort.Slice(sortedCharacters, func(i, j int) bool {
		return sortedCharacters[i].Name > sortedCharacters[j].Name
	})
	return sortedCharacters
}

func filterCharactersByFemale(AllCharacters []service.Characters) []service.Characters {
	var females []service.Characters
	for _, character := range AllCharacters {
		if character.Gender == "Female" {
			females = append(females, character)
		}
	}
	return females
}

func filterCharactersByMale(AllCharacters []service.Characters) []service.Characters {
	var males []service.Characters
	for _, character := range AllCharacters {
		if character.Gender == "Male" {
			males = append(males, character)
		}
	}
	return males
}

func filterCharactersByNation(names []string, AllCharacters []service.Characters) []service.Characters {
	var nations []service.Characters
	for _, nation := range names {
		for _, characters := range AllCharacters {
			if characters.Nation == nation {
				nations = append(nations, characters)
			}
		}
	}
	return nations
}

func filterCharactersByRarity(rarity int, AllCharacters []service.Characters) []service.Characters {
	var Characters []service.Characters
	for _, characters := range AllCharacters {
		if characters.Rarity == rarity {
			Characters = append(Characters, characters)
		}
	}
	return Characters
}

func filterCharactersByVision(visions []string, AllCharacters []service.Characters) []service.Characters {
	var Characters []service.Characters
	for _, vision := range visions {
		for _, character := range AllCharacters {
			if character.Vision == vision {
				Characters = append(Characters, character)
			}
		}
	}
	return Characters
}

func sortCharactersRecent(characters []service.Characters) []service.Characters {
	sort.Slice(characters, func(i, j int) bool {
		date1, err1 := time.Parse("2006-01-02", characters[i].Release)
		date2, err2 := time.Parse("2006-01-02", characters[j].Release)
		if err1 != nil || err2 != nil {
			return false
		}
		return date1.After(date2)
	})
	return characters
}

func sortCharactersByOldest(characters []service.Characters) []service.Characters {
	sort.Slice(characters, func(i, j int) bool {
		date1, err1 := time.Parse("2006-01-02", characters[i].Release)
		date2, err2 := time.Parse("2006-01-02", characters[j].Release)
		if err1 != nil || err2 != nil {
			return false
		}
		return date1.Before(date2)
	})
	return characters
}

func CharacterFiltersChecker(filters CharacterFilters) bool {
	if len(filters.Gender) > 2 {
		return false
	}
	for _, gender := range filters.Gender {
		if !(gender == "" || gender == "male" || gender == "female") {
			return false
		}
	}
	if len(filters.Nations) > 7 {
		return false
	}
	for _, nation := range filters.Nations {
		if !(nation == "" || nation == "Mondstadt" || nation == "Fontaine" || nation == "Inazuma" || nation == "Liyue" || nation == "Natlan" || nation == "Sumeru" || nation == "Snezhnaya") {
			return false
		}
	}
	if len(filters.Rarity) > 2 {
		return false
	}
	for _, rarity := range filters.Rarity {
		if !(rarity == "" || rarity == "5" || rarity == "4") {
			return false
		}
	}
	if len(filters.Element) > 7 {
		return false
	}
	for _, element := range filters.Element {
		if !(element == "" || element == "Geo" || element == "Electro" || element == "Hydro" || element == "Pyro" || element == "Dendro" || element == "Anemo" || element == "Cryo") {
			return false
		}
	}
	return (filters.ReleaseDateFilter == "" || filters.ReleaseDateFilter == "az" || filters.ReleaseDateFilter == "za" || filters.ReleaseDateFilter == "recent" || filters.ReleaseDateFilter == "oldest")
}
