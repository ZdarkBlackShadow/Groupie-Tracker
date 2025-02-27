package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = append(body[:len(body) - 16], body[len(body) - 1:]...)
		errDecode := json.Unmarshal(body, &decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return map[string]Potion{}
}

func GetAllPotions() ([]Potion, error) {
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
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var decodeData []string
		errDecode := json.Unmarshal(body, &decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		L := []Potion{}
		for i, name := range decodeData {
			L = append(L, m[name])
			L[i].ImageUrl = "https://genshin.jmp.blue/consumables/potions/" + name
			L[i].Id = name
		}
		return L, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []Potion{}, nil
}

func GetDetailsAboutOnePotion(id string, AllPotions []Potion) Potion {
	for _, potion := range AllPotions {
		if potion.Id == id {
			return potion
		}
	}
	return Potion{}
}
