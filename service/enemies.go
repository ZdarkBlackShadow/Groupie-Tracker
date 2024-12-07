package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneEnemie(name string) Enemies {
	urlApi := "https://genshin.jmp.blue/enemies/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("Erreur lors de la requette http : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		fmt.Printf("Erreur lors de la requete client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData Enemies
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return Enemies{}
}

func GetNamesOfAllTheEnemies() []string {
	urlApi := "https://genshin.jmp.blue/enemies"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("Error during http request : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		fmt.Printf("Error during client request : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("error during decoding : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Code error : %v, error message : %v", res.StatusCode, res.Status)
	}
	return []string{}
}

func GetAllEnemiesDetails() []Enemies {
	AllNames := GetNamesOfAllTheEnemies()
	AllBossDetails := []Enemies{}
	for _, name := range AllNames {
		AllBossDetails = append(AllBossDetails, GetAllDetailsAboutOneEnemie(name))
	}
	return AllBossDetails
}
