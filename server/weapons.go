package server

import (
	"log"
	"net/http"
	"sort"
	"strconv"

	"main.go/service"
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

type WeaponFilters struct {
	Rarity     []string
	Type       []string
	BaseAttack string
	SortBy     string
}

func PaginateWeapons(data []service.Weapon, page, itemsPerPage int) ([]service.Weapon, int) {
	totalItems := len(data)
	totalPages := totalItems / itemsPerPage

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

func GetWeaponFilters(r *http.Request) WeaponFilters {
	filters := WeaponFilters{
		Rarity:     r.URL.Query()["rarity_filter"],
		Type:       r.URL.Query()["type_filter"],
		BaseAttack: r.URL.Query().Get("attack_filter"),
		SortBy:     r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplyWeaponFilters(filters WeaponFilters) []service.Weapon {
	var weapons []service.Weapon
	if len(filters.Rarity) > 0 {
		weapons = SortWeaponsByRarity(filters.Rarity, API_Data.AllWeapons)
	} else {
		weapons = API_Data.AllWeapons
	}
	if len(filters.Type) > 0 {
		weapons = SortWeaponsByType(filters.Type, weapons)
	}
	if filters.BaseAttack != "" {
		if filters.BaseAttack == "highest" {
			weapons = SortWeaponsByHighestAttack(weapons)
		} else if filters.BaseAttack == "lowest" {
			weapons = SortWeaponsByLowestAttack(weapons)
		}
	}
	if filters.SortBy != "" {
		if filters.SortBy == "az" {
			weapons = SortWeaponsByAZ(weapons)
		} else if filters.SortBy == "za" {
			weapons = SortWeaponsByZA(weapons)
		}
	}
	return weapons
}

func Weapons(w http.ResponseWriter, r *http.Request) {
	AllWeopns := ApplyWeaponFilters(GetWeaponFilters(r))

	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}
	pagedData, totalPages := PaginateWeapons(AllWeopns, page, 20)
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
func WeaponDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := DataWeaponDetails{
		Data:    service.GetAllDataAboutOneWeapon(Id),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "weaponsDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func SortWeaponsByRarity(raritys []string, weapon []service.Weapon) []service.Weapon {
	weapons := []service.Weapon{}
	for _, rarity := range raritys {
		for _, weapon := range weapon {
			intRarity, _ := strconv.Atoi(rarity)
			if weapon.Rarity == intRarity {
				weapons = append(weapons, weapon)
			}
		}
	}
	return weapons
}

func SortWeaponsByType(types []string, weapon []service.Weapon) []service.Weapon {
	weapons := []service.Weapon{}
	for _, t := range types {
		for _, weapon := range weapon {
			if weapon.Type == t {
				weapons = append(weapons, weapon)
			}
		}
	}
	return weapons
}

func SortWeaponsByHighestAttack(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].BaseAttack > sortedWeapons[j].BaseAttack
	})
	return sortedWeapons
}

func SortWeaponsByLowestAttack(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name < sortedWeapons[j].Name
	})
	return sortedWeapons
}

func SortWeaponsByAZ(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name < sortedWeapons[j].Name
	})
	return sortedWeapons
}

func SortWeaponsByZA(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name > sortedWeapons[j].Name
	})
	return sortedWeapons
}
