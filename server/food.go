package server

import (
	"math"
	"net/http"
	"strconv"

	"main.go/service"
)

type FoodStructData struct {
	Data        []service.FoodStruct
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type FoodDetailStructData struct {
	Data    service.FoodStruct
	IsLogin bool
}

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

func ApplyFoodFilters(filters FoodFilters) []service.FoodStruct {
	Foods := []service.FoodStruct{}
	if filters.Recipe != "" {
		if filters.Recipe == "1" {
			for _, food := range API_Data.AllFood {
				if food.HasRecipe {
					Foods = append(Foods, food)
				}
			}
		} else {
			for _, food := range API_Data.AllFood {
				if !food.HasRecipe {
					Foods = append(Foods, food)
				}
			}
		}
	}
	if len(filters.Type) > 0 {
		for _, food := range API_Data.AllFood {
			if Contains(filters.Type, food.Type) {
				Foods = append(Foods, food)
			}
		}
	}
	if len(filters.Rarity) > 0 {
		for _, food := range API_Data.AllFood {
			if Contains(filters.Rarity, strconv.Itoa(food.Rarity)) {
				Foods = append(Foods, food)
			}
		}
	}
	if len(filters.Sort) > 0 {
		if filters.Sort == "az" {
			Foods = SortFoodsAZ(Foods)
		} else if filters.Sort == "za" {
			Foods = SortFoodsZA(Foods)
		}
	}
	if len(Foods) == 0 {
		Foods = SortFoodsAZ(API_Data.AllFood)
	} else {
		Foods = RemoveDuplicatesFood(Foods)
	}
	return Foods
}

func Food(w http.ResponseWriter, r *http.Request) {
	filters := GetFoodFilters(r)
	DataFood := ApplyFoodFilters(filters)
	pageParam := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	ItemsPerPage := 16
	Data, totalPages := PaginateFood(DataFood, page, ItemsPerPage)
	DataForTheTemplate := FoodStructData{
		Data:        Data,
		TotalPages:  totalPages,
		CurrentPage: page,
		IsLogin:     IsLogin,
	}
	err = Templates.ExecuteTemplate(w, "food", DataForTheTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func FoodDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	Data := FoodDetailStructData{
		Data:    service.GetDetailsOfFood(id, API_Data.AllFood),
		IsLogin: IsLogin,
	}
	err := Templates.ExecuteTemplate(w, "foodDetails", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RemoveDuplicatesFood(foods []service.FoodStruct) []service.FoodStruct {
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

func SortFoodsAZ(foods []service.FoodStruct) []service.FoodStruct {
	for i := 0; i < len(foods); i++ {
		for j := i + 1; j < len(foods); j++ {
			if foods[i].Name > foods[j].Name {
				foods[i], foods[j] = foods[j], foods[i]
			}
		}
	}
	return foods
}

func SortFoodsZA(foods []service.FoodStruct) []service.FoodStruct {
	for i := 0; i < len(foods); i++ {
		for j := i + 1; j < len(foods); j++ {
			if foods[i].Name < foods[j].Name {
				foods[i], foods[j] = foods[j], foods[i]
			}
		}
	}
	return foods
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
