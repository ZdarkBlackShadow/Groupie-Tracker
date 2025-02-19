package utils

import (
	"log"
	"sync"

	"main.go/service"
)

var (
	Mu       sync.Mutex
	Progress int
)

type AllData struct {
	ALLArtifacts  []service.ArtifactDetails
	AllBoss       []service.BossStruct
	AllCharacters []service.Characters
	AllDomain     []service.Domain
	AllElement    []service.Element
	AllEnnemies   []service.Enemies
	AllWeapons    []service.Weapon
	AllPotions    []service.Potion
	AllFood       []service.FoodStruct
}

func (API_Data *AllData) LoadAllData() {
	var err error
	Progress = 0
	//food
	food := service.TransformFoodToSlice(service.GetAllDetailsOfFood())
	Mu.Lock()
	API_Data.AllFood = food
	Progress = 10
	Mu.Unlock()
	//artifacts
	artifacts, err := service.GetAllArtifactsDetails()
	if err != nil {
		log.Fatal(err)
	}
	Mu.Lock()
	API_Data.ALLArtifacts = artifacts
	Progress = 20
	Mu.Unlock()
	//potions
	potions := service.GetAllPotions()
	Mu.Lock()
	API_Data.AllPotions = potions
	Progress = 30
	Mu.Unlock()
	//boss
	bosses, err := service.GetAllBossDetails()
	if err != nil {
		log.Fatal(err)
	}
	Mu.Lock()
	API_Data.AllBoss = bosses
	Progress = 40
	Mu.Unlock()
	//characters
	characters, err := service.GetAllCharactersDetails()
	if err != nil {
		log.Fatal(err)
	}
	Mu.Lock()
	API_Data.AllCharacters = characters
	Progress = 50
	Mu.Unlock()
	//Domain
	domains, err := service.GetAllDomainsDetails()
	if err != nil {
		log.Fatal(err)
	}
	Mu.Lock()
	API_Data.AllDomain = domains
	Progress = 60
	Mu.Unlock()
	//Elements
	elements := service.GetAllElementsDetails()
	Mu.Lock()
	API_Data.AllElement = elements
	Progress = 70
	Mu.Unlock()
	//Enemies
	enemies := service.GetAllEnemiesDetails()
	Mu.Lock()
	API_Data.AllEnnemies = enemies
	Progress = 80
	Mu.Unlock()
	//Weapons
	weapons := service.GetAllWeaponsDetails()
	Mu.Lock()
	API_Data.AllWeapons = weapons
	Progress = 90
	Mu.Unlock()
	Mu.Lock()
	Progress = 100
	Mu.Unlock()
}
