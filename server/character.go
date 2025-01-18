package server

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"main.go/service"
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

type CharacterFilters struct {
	Gender            []string
	Nations           []string
	Rarity            []string
	Element           []string
	ReleaseDateFilter string
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

func GetCharacterFilters(r *http.Request) CharacterFilters {
	filters := CharacterFilters{
		Gender:            r.URL.Query()["gender_filter"],
		Nations:           r.URL.Query()["nations_filter"],
		Rarity:            r.URL.Query()["rarity_filter"],
		Element:           r.URL.Query()["elements_filter"],
		ReleaseDateFilter: r.URL.Query().Get("release_date_filter"),
	}
	return filters
}

func ApplyCharacterFilters(filters CharacterFilters) []service.Characters {
	Characters := []service.Characters{}
	if len(filters.Gender) == 1 {
		if filters.Gender[0] == "male" {
			MaleCharacters := FilterCharactersByMale()
			Characters = append(Characters, MaleCharacters...)
		} else {
			FemaleCharacters := FilterCharactersByFemale()
			Characters = append(Characters, FemaleCharacters...)
		}
	}
	if len(filters.Nations) > 0 {
		NationCharacters := FilterCharactersByNation(filters.Nations)
		Characters = append(Characters, NationCharacters...)
	}
	if len(filters.Rarity) == 1 {
		for _, rarity := range filters.Rarity {
			rarityInt, err := strconv.Atoi(rarity)
			if err != nil {
				continue
			}
			RarityCharacters := FilterCharactersByRarity(rarityInt)
			Characters = append(Characters, RarityCharacters...)
		}
	}
	if len(filters.Element) > 0 {
		VisionCharacters := FilterCharactersByVision(filters.Element)
		Characters = append(Characters, VisionCharacters...)
	}
	if len(Characters) == 0 {
		Characters = API_Data.AllCharacters
	} else {
		Characters = RemoveDuplicates(Characters)
	}
	if filters.ReleaseDateFilter != "" {
		if filters.ReleaseDateFilter == "recent" {
			Characters = SortCharactersRecent(Characters)
		} else if filters.ReleaseDateFilter == "oldest" {
			Characters = SortCharactersByOldest(Characters)
		} else if filters.ReleaseDateFilter == "az" {
			Characters = FilterCharactersByAZ(Characters)
		} else if filters.ReleaseDateFilter == "za" {
			Characters = FilterCharactersByZA(Characters)
		}
	}
	return Characters
}

func RemoveDuplicates(characters []service.Characters) []service.Characters {
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

func Characters(w http.ResponseWriter, r *http.Request) {
	allCharacters := ApplyCharacterFilters(GetCharacterFilters(r))
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
		IsLogin:     IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "characters", dataArtifacts)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func CharactersDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Datat := DataCharactersDetails{
		Data:    service.GetAllDetailsAboutOneCharacters(Id),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "characterDetails", Datat)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func FilterCharactersByAZ(characters []service.Characters) []service.Characters {
	sortedCharacters := make([]service.Characters, len(characters))
	copy(sortedCharacters, characters)

	sort.Slice(sortedCharacters, func(i, j int) bool {
		return sortedCharacters[i].Name < sortedCharacters[j].Name
	})
	return sortedCharacters
}

func FilterCharactersByZA(characters []service.Characters) []service.Characters {
	sortedCharacters := make([]service.Characters, len(characters))
	copy(sortedCharacters, characters)

	sort.Slice(sortedCharacters, func(i, j int) bool {
		return sortedCharacters[i].Name > sortedCharacters[j].Name
	})
	return sortedCharacters
}

func FilterCharactersByFemale() []service.Characters {
	var females []service.Characters
	for _, character := range API_Data.AllCharacters {
		if character.Gender == "Female" {
			females = append(females, character)
		}
	}
	return females
}

func FilterCharactersByMale() []service.Characters {
	var males []service.Characters
	for _, character := range API_Data.AllCharacters {
		if character.Gender == "Male" {
			males = append(males, character)
		}
	}
	return males
}

func FilterCharactersByNation(names []string) []service.Characters {
	var nations []service.Characters
	for _, nation := range names {
		for _, characters := range API_Data.AllCharacters {
			if characters.Nation == nation {
				nations = append(nations, characters)
			}
		}
	}
	return nations
}

func FilterCharactersByRarity(rarity int) []service.Characters {
	var Characters []service.Characters
	for _, characters := range API_Data.AllCharacters {
		if characters.Rarity == rarity {
			Characters = append(Characters, characters)
		}
	}
	return Characters
}

func FilterCharactersByVision(visions []string) []service.Characters {
	var Characters []service.Characters
	for _, vision := range visions {
		for _, character := range API_Data.AllCharacters {
			if character.Vision == vision {
				Characters = append(Characters, character)
			}
		}
	}
	return Characters
}

func SortCharactersRecent(characters []service.Characters) []service.Characters {
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

func SortCharactersByOldest(characters []service.Characters) []service.Characters {
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
