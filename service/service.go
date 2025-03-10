package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Service() {
	fmt.Println("service")
	//lien de l'API https://genshin.jmp.blue/
}

func GetListOfAvailableImage(Url string) []string {
	urlApi := Url + "/list"
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
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		result := []string{}
		for _, element := range decodeData {
			result = append(result, Url+"/"+element)
		}
		return result
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []string{}
}
