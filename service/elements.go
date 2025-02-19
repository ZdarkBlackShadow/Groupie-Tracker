package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneElements(name string) (Element, error) { //rajouter un int pour le code d'erreur
	urlApi := "https://genshin.jmp.blue/elements/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return Element{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return Element{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData Element
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			return Element{}, errDecode
		}
		decodeData.ImageUrl = urlApi + "/icon"
		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return Element{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetNamesOfAllTheElements() ([]string, error) {
	urlApi := "https://genshin.jmp.blue/elements"
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

func GetAllElementsDetails() []Element {
	var err error
	AllNames, err := GetNamesOfAllTheElements()
	if err != nil {

	}
	AllBossDetails := []Element{}
	for _, name := range AllNames {
		element, err := GetAllDetailsAboutOneElements(name)
		if err != nil {
			log.Fatal(err)
		}
		AllBossDetails = append(AllBossDetails, element)
	}
	return AllBossDetails
}
