package utils

import (
	"math"
	"net/http"
	"strconv"

	"main.go/service"
)

type FoodFilters struct {
	Recipe string
	Type   []string
	Rarity []string
	Sort   string
}

func PaginateFood(data []service.FoodStruct, page, itemsPerPage int) ([]service.FoodStruct, int) {
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

func GetFoodFilters(r *http.Request) FoodFilters {
	filters := FoodFilters{
		Recipe: r.URL.Query().Get("recipe_filter"),
		Type:   r.URL.Query()["types_filter"],
		Rarity: r.URL.Query()["rarity_filter"],
		Sort:   r.URL.Query().Get("sort_filter"),
	}
	return filters
}

func ApplyFoodFilters(filters FoodFilters, AllFoods []service.FoodStruct) []service.FoodStruct {
	Foods := []service.FoodStruct{}
	if filters.Recipe != "" {
		if filters.Recipe == "1" {
			for _, food := range AllFoods {
				if food.HasRecipe {
					Foods = append(Foods, food)
				}
			}
		} else {
			for _, food := range AllFoods {
				if !food.HasRecipe {
					Foods = append(Foods, food)
				}
			}
		}
	}
	if len(filters.Type) > 0 {
		for _, food := range AllFoods {
			if contains(filters.Type, food.Type) {
				Foods = append(Foods, food)
			}
		}
	}
	if len(filters.Rarity) > 0 {
		for _, food := range AllFoods {
			if contains(filters.Rarity, strconv.Itoa(food.Rarity)) {
				Foods = append(Foods, food)
			}
		}
	}
	if len(filters.Sort) > 0 {
		if filters.Sort == "az" {
			Foods = sortFoodsAZ(Foods)
		} else if filters.Sort == "za" {
			Foods = sortFoodsZA(Foods)
		}
	}
	if len(Foods) == 0 {
		Foods = sortFoodsAZ(AllFoods)
	} else {
		Foods = removeDuplicatesFood(Foods)
	}
	return Foods
}


func removeDuplicatesFood(foods []service.FoodStruct) []service.FoodStruct {
	encountered := map[string]bool{}
	result := []service.FoodStruct{}

	for _, food := range foods {
		if !encountered[food.Id] {
			encountered[food.Id] = true
			result = append(result, food)
		}
	}
	return result
}

func sortFoodsAZ(foods []service.FoodStruct) []service.FoodStruct {
	for i := 0; i < len(foods); i++ {
		for j := i + 1; j < len(foods); j++ {
			if foods[i].Name > foods[j].Name {
				foods[i], foods[j] = foods[j], foods[i]
			}
		}
	}
	return foods
}

func sortFoodsZA(foods []service.FoodStruct) []service.FoodStruct {
	for i := 0; i < len(foods); i++ {
		for j := i + 1; j < len(foods); j++ {
			if foods[i].Name < foods[j].Name {
				foods[i], foods[j] = foods[j], foods[i]
			}
		}
	}
	return foods
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
