package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

func GetAllNameOfArtifacts() []string {
	urlApi := "https://genshin.jmp.blue/artifacts"
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
		fmt.Printf("Erreur lors de la requete client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []string{}
}

func GetAllDataAboutOneArtifact(name string) ArtifactDetails {
	urlApi := "https://genshin.jmp.blue/artifacts/" + name
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
		var decodeData ArtifactDetails
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("erreur lors du decodage : %v\n", errDecode)
		}
		if name == "glacier-and-snowfield" {
			decodeData.ImageURL = "/static/image/Exception/Glacier and Snowfield/flower.webp"
			decodeData.AllUrlImageAvailable = []string{"/static/image/Exception/Glacier and Snowfield/flower.webp"}
		} else {
			AvailableImage := GetListOfAvailableImage(urlApi)
			decodeData.ImageURL = AvailableImage[0]
			decodeData.AllUrlImageAvailable = AvailableImage
		}

		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}

	return ArtifactDetails{}
}
func GetAllArtifactsDetails() []ArtifactDetails {
	AllNames := GetAllNameOfArtifacts()
	AllArtifactsDetails := []ArtifactDetails{}
	for _, name := range AllNames {
		if name != "prayers-to-the-firmament" {
			AllArtifactsDetails = append(AllArtifactsDetails, GetAllDataAboutOneArtifact(name))
		}
	}
	return AllArtifactsDetails
}

func ArtifactsFilterByName(data []ArtifactDetails) []ArtifactDetails {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Name < data[j].Name
	})
	return data
}

func ArtifactsFilterByRarity(data []ArtifactDetails) []ArtifactDetails {
	sort.Slice(data, func(i, j int) bool {
		return data[i].MaxRarity < data[j].MaxRarity
	})
	return data
}
