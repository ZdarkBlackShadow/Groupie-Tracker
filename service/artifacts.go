package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getAllNameOfArtifacts() ([]string, error) {
	var (
		urlApi     string
		httpClient http.Client
		req        *http.Request
		res        *http.Response
		decodeData []string
		err        error
	)
	urlApi = "https://genshin.jmp.blue/artifacts"
	httpClient = http.Client{
		Timeout: time.Second * 5,
	}
	req, err = http.NewRequest(http.MethodGet, urlApi, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Set("User-Agent", "Ynov Campus module groupie tracker")
	res, err = httpClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	if res.StatusCode == http.StatusOK {
		err = json.NewDecoder(res.Body).Decode(&decodeData)
		if err != nil {
			fmt.Printf("erreur lors du decodage : %v\n", err)
		}
		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []string{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetAllDataAboutOneArtifact(name string) (ArtifactDetails, error) {
	urlApi := "https://genshin.jmp.blue/artifacts/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return ArtifactDetails{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return ArtifactDetails{}, Errres
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

		return decodeData, nil
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}

	return ArtifactDetails{}, errors.New("Status code is not 200 but " + res.Status)
}
func GetAllArtifactsDetails() ([]ArtifactDetails, error) {
	AllNames, err := getAllNameOfArtifacts()
	if err != nil {
		return []ArtifactDetails{}, err
	}
	AllArtifactsDetails := []ArtifactDetails{}
	for _, name := range AllNames {
		if name != "prayers-to-the-firmament" {
			artifact, err := GetAllDataAboutOneArtifact(name)
			if err != nil {
				log.Fatal(err)
			}
			AllArtifactsDetails = append(AllArtifactsDetails, artifact)
		}
	}
	return AllArtifactsDetails, nil
}
