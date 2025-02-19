package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		image, err := checkResponseImageForEnemies(urlApi + "/list")
		if err != nil {
			log.Fatal(err) 
		}
		if image {
			decodeData.ImageUrl = urlApi + "/icon"
		} else {
			decodeData.ImageUrl = "/static/image/NoImageAvaliable.webp"
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

func checkResponseImageForEnemies(apiURL string) (bool, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var responseArray []string
	if json.Unmarshal(body, &responseArray) == nil {
		if len(responseArray) == 2 && responseArray[0] == "icon" && responseArray[1] == "portrait" {
			return true, nil
		}
	}

	var responseMap map[string]string
	if json.Unmarshal(body, &responseMap) == nil {
		if val, exists := responseMap["error"]; exists && val == "No images for enemies/consecrated-beast exist" {
			return false, nil
		}
	}

	return false, nil
}
