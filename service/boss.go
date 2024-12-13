package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneBoss(name string) BossStruct {
	urlApi := "https://genshin.jmp.blue/boss/weekly-boss/" + name
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
		var decodeData BossStruct
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		decodeData.ImageURL = "https://genshin.jmp.blue/boss/weekly-boss/" + name + "/portrait"
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return BossStruct{}
}

func GetNamesOfAllTheBosses() []string {
	urlApi := "https://genshin.jmp.blue/boss/weekly-boss"
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

func GetAllBossDetails() []BossStruct {
	AllNames := GetNamesOfAllTheBosses()
	AllBossDetails := []BossStruct{}
	for _, name := range AllNames {
		AllBossDetails = append(AllBossDetails, GetAllDetailsAboutOneBoss(name))
	}
	return AllBossDetails
}
