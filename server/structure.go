package server

import (
	"main.go/service"
)

type Register struct {
	Email        string
	Password     string
	Collecttions Collecttions
}

type Collecttions struct {
	Artifacts  []service.ArtifactDetails
	Boss       []service.BossStruct
	Characters []service.Characters
	Domains    []service.Domain
	Elements   []service.Element
	Enemies    []service.Enemies
	Weapons    []service.Weapon
	Materials  CollenctionsMaterials
}

type CollenctionsMaterials struct {
	BossMaterials       []service.BossMaterial
	CharacterAscensions []service.CharactersAscensionDataStructure
	CharacterExperience []service.CharacterExperienceDataStruct
	CommonAscensions    []service.CommonAscensionDetails
	Ingredients         []service.Ingredient
	LocalSpecialities   []service.Specialty
	//une autre sructure avant
	BossTalents          []service.Talent
	WheaponAscensionSets []service.WeaponAscensionSet
	WeaponsExperiences   []service.WeaponAscensionItem
	Foods                []service.FoodStruct
	Potions              []service.Potion
}
