package utils

import "main.go/service"

func HasIdInEnemie(id string, AllEnemies []service.Enemies) bool {
	for _, enemie := range AllEnemies {
		if id == enemie.Id {
			return true
		}
	}
	return false
}