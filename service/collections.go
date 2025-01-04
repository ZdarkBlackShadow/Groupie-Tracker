package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetCollections() Register {
	file, err := os.Open("./data/data.json")
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		return Register{}
	}
	defer file.Close()

	// Lecture du contenu du fichier
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		return Register{}
	}

	// Déclaration d'une structure générique pour contenir les données JSON
	var data Register

	// Décodage des données JSON
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Printf("Erreur lors du décodage JSON: %v\n", err)
		return Register{}
	}
	return data
}
