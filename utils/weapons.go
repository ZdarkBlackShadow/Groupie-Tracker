package utils

import (
	"math"
	"net/http"
	"sort"
	"strconv"

	"main.go/service"
)

type WeaponFilters struct {
	Rarity     []string
	Type       []string
	BaseAttack string
	SortBy     string
}

func PaginateWeapons(data []service.Weapon, page, itemsPerPage int) ([]service.Weapon, int, int) {
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

func GetWeaponFilters(r *http.Request) WeaponFilters {
	filters := WeaponFilters{
		Rarity:     r.URL.Query()["rarity_filter"],
		Type:       r.URL.Query()["type_filter"],
		BaseAttack: r.URL.Query().Get("attack_filter"),
		SortBy:     r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplyWeaponFilters(filters WeaponFilters, AllWeapons []service.Weapon) []service.Weapon {
	var weapons []service.Weapon
	if len(filters.Rarity) > 0 {
		weapons = sortWeaponsByRarity(filters.Rarity, AllWeapons)
	} else {
		weapons = AllWeapons
	}
	if len(filters.Type) > 0 {
		weapons = sortWeaponsByType(filters.Type, weapons)
	}
	if filters.BaseAttack != "" {
		if filters.BaseAttack == "highest" {
			weapons = sortWeaponsByHighestAttack(weapons)
		} else if filters.BaseAttack == "lowest" {
			weapons = sortWeaponsByLowestAttack(weapons)
		}
	}
	if filters.SortBy != "" {
		if filters.SortBy == "az" {
			weapons = sortWeaponsByAZ(weapons)
		} else if filters.SortBy == "za" {
			weapons = sortWeaponsByZA(weapons)
		}
	}
	return weapons
}

func sortWeaponsByRarity(raritys []string, weapon []service.Weapon) []service.Weapon {
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

func sortWeaponsByType(types []string, weapon []service.Weapon) []service.Weapon {
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

func sortWeaponsByHighestAttack(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].BaseAttack > sortedWeapons[j].BaseAttack
	})
	return sortedWeapons
}

func sortWeaponsByLowestAttack(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name < sortedWeapons[j].Name
	})
	return sortedWeapons
}

func sortWeaponsByAZ(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name < sortedWeapons[j].Name
	})
	return sortedWeapons
}

func sortWeaponsByZA(weapons []service.Weapon) []service.Weapon {
	sortedWeapons := make([]service.Weapon, len(weapons))
	copy(sortedWeapons, weapons)

	sort.Slice(sortedWeapons, func(i, j int) bool {
		return sortedWeapons[i].Name > sortedWeapons[j].Name
	})
	return sortedWeapons
}

func WeaponsFiltersChecker(filters WeaponFilters) bool {
	if len(filters.Rarity) > 3 {
		return false
	}
	for _, rarity := range filters.Rarity {
		if !(rarity == "" || rarity == "3" || rarity == "4" || rarity == "5") {
			return false
		}
	}
	if len(filters.Type) > 5 {
		return false
	}
	for _, WeaponType := range filters.Type {
		if !(WeaponType == "" || WeaponType == "Sword" || WeaponType == "Polearm" || WeaponType == "Bow" || WeaponType == "Claymore" || WeaponType == "Catalyst") {
			return false
		}
	}
	if !(filters.SortBy == "" || filters.SortBy == "az" || filters.SortBy == "za") {
		return false
	}
	return (filters.BaseAttack == "" || filters.BaseAttack == "highest" || filters.BaseAttack == "weakest")
}
