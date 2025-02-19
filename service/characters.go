package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneCharacters(name string) (Characters, error) {
	urlApi := "https://genshin.jmp.blue/characters/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return Characters{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return Characters{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData Characters
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		decodeData.ImageUrl = urlApi + "/card"
		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return Characters{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetNamesOfAllTheCharacters() ([]string, error) {
	urlApi := "https://genshin.jmp.blue/characters"
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

func GetAllCharactersDetails() ([]Characters, error) {
	var err error
	AllNames, err := GetNamesOfAllTheCharacters()
	if err != nil {
		return []Characters{}, err
	}
	AllBossDetails := []Characters{}
	for _, name := range AllNames {
		boss, err := GetAllDetailsAboutOneCharacters(name)
		if err != nil {
			return []Characters{}, err
		}
		AllBossDetails = append(AllBossDetails, boss)
	}
	return AllBossDetails, nil
}
