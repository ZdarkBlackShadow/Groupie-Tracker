package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Récupère tous les noms d'armes
func GetAllNameOfWeapon() []string {
	urlApi := "https://genshin.jmp.blue/weapons" // url de l'api
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("Erreur lors de la requête http : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		fmt.Printf("Erreur lors de la requête client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("Erreur lors du décodage : %v\n", errDecode)
		}
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}
	return []string{}
}

// Récupère toutes les données sur une arme spécifique
func GetAllDataAboutOneWeapon(name string) Weapon {
	urlApi := "https://genshin.jmp.blue/weapons/" + name // url de l'api
	httpClient := http.Client{
		Timeout: time.Second * 5, // timeout de 5 secondes, pour prévoir le temps de chargement
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Printf("Erreur lors de la requête http : %v\n", errReq)
	}
	req.Header.Set("User-Agent", "Ynov Campus") //optionnel
	res, Errres := httpClient.Do(req)           // envoie de la requête
	if Errres != nil {
		fmt.Printf("Erreur lors de la requête client : %v\n", Errres)
	}
	if res.StatusCode == http.StatusOK { // si la requête est un succès
		var decodeData Weapon                                      // déclaration de la variable qui va contenir les données
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData) // décodage des données
		if errDecode != nil {
			fmt.Printf("Erreur lors du décodage : %v\n", errDecode)
		}
		decodeData.Id = name
		decodeData.ImageUrl = urlApi + "/icon" // ajout de l'url de l'image
		return decodeData
	} else {
		fmt.Printf("Erreur code : %v, erreur message : %v", res.StatusCode, res.Status)
	}

	return Weapon{}
}

func GetAllWeaponsDetails() []Weapon {
	AllNames := GetAllNameOfWeapon()
	AllWeoponDetails := []Weapon{}
	for _, name := range AllNames {
		AllWeoponDetails = append(AllWeoponDetails, GetAllDataAboutOneWeapon(name))
	}
	return AllWeoponDetails
}
