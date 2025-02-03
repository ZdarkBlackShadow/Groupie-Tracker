package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetAllPotionsMap() map[string]Potion {
	UrlAPI := "https://genshin.jmp.blue/consumables/potions"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, UrlAPI, nil)
	if errReq != nil {
		fmt.Printf("Erreur lors de la requette http : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		fmt.Printf("Erreur lors de la requete client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData map[string]Potion
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return map[string]Potion{}
}

func GetAllPotions() []Potion {
	m := GetAllPotionsMap()
	UrlAPI := "https://genshin.jmp.blue/consumables/potions/list"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, UrlAPI, nil)
	if errReq != nil {
		fmt.Printf("Erreur lors de la requette http : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		fmt.Printf("Erreur lors de la requete client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		L := []Potion{}
		for i, name := range decodeData {
			L = append(L, m[name])
			L[i].ImageUrl = "https://genshin.jmp.blue/consumables/potions/" + name
			L[i].Id = name
		}
		return L
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []Potion{}
}

func GetDetailsAboutOnePotion(id string, AllPotions []Potion) Potion {
	for _, potion := range AllPotions {
		if potion.Id == id {
			return potion
		}
	}
	return Potion{}
}
