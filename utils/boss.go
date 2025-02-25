package utils

import "main.go/service"

//Ficher pour verifie l'id des boss

func HasIdInBoss(id string, AllBoss []service.BossStruct) bool {
	for _, boss := range AllBoss {
		if boss.Id == id {
			return true
		}
	}
	return false
}