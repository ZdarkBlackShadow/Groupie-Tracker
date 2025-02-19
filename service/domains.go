package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneDomains(name string) (Domain, error) {
	urlApi := "https://genshin.jmp.blue/domains/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return Domain{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return Domain{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData Domain
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return Domain{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetNamesOfAllTheDomains() ([]string, error) {
	urlApi := "https://genshin.jmp.blue/domains"
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
	}
	return []string{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetAllDomainsDetails() ([]Domain, error) {
	var err error
	AllNames, err := GetNamesOfAllTheDomains()
	if err != nil {
		return []Domain{}, err
	}
	AllBossDetails := []Domain{}
	for _, name := range AllNames {
		domain, err := GetAllDetailsAboutOneDomains(name)
		if err != nil {
			return []Domain{}, err
		}
		AllBossDetails = append(AllBossDetails, domain)
	}
	return AllBossDetails, nil
}
