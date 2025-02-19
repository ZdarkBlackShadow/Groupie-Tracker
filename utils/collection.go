package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"main.go/service"
)

//rajouter une gestion d'erreur

func GetCollectionOfTheActualUser(ActualUser service.Register) service.Register {
	file, err := os.Open("./data/data.json")
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		return service.Register{}
	}
	defer file.Close()

	// Lecture du contenu du fichier
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier: %v\n", err)
		return service.Register{}
	}

	// Déclaration d'une structure générique pour contenir les données JSON
	var data []service.Register

	// Décodage des données JSON
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Printf("Erreur lors du décodage JSON: %v\n", err)
		return service.Register{}
	}
	for _, register := range data {
		if register.Email == ActualUser.Email {
			return register
		}
	}
	return service.Register{}
}

func RemoveFromTheCollection(ElementToRemove service.CollectionsStruct, ActualUser *service.Register) {
	for i, element := range ActualUser.Collecttions.Collections {
		if element.Name == ElementToRemove.Name && element.Type == ElementToRemove.Type {
			ActualUser.Collecttions.Collections = append(ActualUser.Collecttions.Collections[:i], ActualUser.Collecttions.Collections[i+1:]...)
			break
		}
	}
	var data []service.Register
	filecontent, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
	}
	if err := json.Unmarshal(filecontent, &data); err != nil {
		log.Fatalf("Erreur lors du décodage JSON : %v", err)
	}
	for i, element := range data {
		if element.Email == ActualUser.Email && element.Password == ActualUser.Password {
			data[i].Collecttions = ActualUser.Collecttions
			break
		}
	}
	updatedJsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Erreur lors de l'encodage JSON : %v", err)
	}
	if err := ioutil.WriteFile("./data/data.json", updatedJsonData, 0644); err != nil {
		log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
	}
}
