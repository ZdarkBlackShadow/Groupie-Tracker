package utils

import "main.go/service"

func HasIdInPotions(id string, AllPotion []service.Potion) bool {
	for _, potion := range AllPotion {
		if potion.Id == id {
			return true
		}
	}
	return false
}
