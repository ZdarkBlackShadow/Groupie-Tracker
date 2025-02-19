package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneBoss(name string) (BossStruct, error) {
	urlApi := "https://genshin.jmp.blue/boss/weekly-boss/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return BossStruct{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return BossStruct{}, Errres
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var decodeData BossStruct
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		decodeData.ImageURL = "https://genshin.jmp.blue/boss/weekly-boss/" + name + "/portrait"
		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return BossStruct{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetNamesOfAllTheBosses() ([]string, error) {
	urlApi := "https://genshin.jmp.blue/boss/weekly-boss"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return []string{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return []string{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("error during decoding : %v\n", errDecode)
		}
		return decodeData, nil
	} else {
		fmt.Printf("Code error : %v, error message : %v", res.StatusCode, res.Status)
	}
	return []string{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetAllBossDetails() ([]BossStruct, error) {
	var err error
	var boss BossStruct
	AllNames, err := GetNamesOfAllTheBosses()
	if err != nil {
		return []BossStruct{}, err
	}
	AllBossDetails := []BossStruct{}
	for _, name := range AllNames {
		boss, err = GetAllDetailsAboutOneBoss(name)
		if err != nil {
			return []BossStruct{}, err
		}
		AllBossDetails = append(AllBossDetails, boss)
	}
	return AllBossDetails, nil
}
