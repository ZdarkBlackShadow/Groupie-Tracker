package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"main.go/service"
)

type ArtifactRequest struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type DataCollectionsStruct struct {
	Data    service.Collecttions
	IsLogin bool
}

func Collections(w http.ResponseWriter, r *http.Request) {
	var data DataCollectionsStruct
	if IsLogin {
		data = DataCollectionsStruct{
			Data:    GetCollectionOfTheActualUser().Collecttions,
			IsLogin: true,
		}
	} else {
		data = DataCollectionsStruct{
			Data:    service.Collecttions{},
			IsLogin: false,
		}
	}
	err := Templates.ExecuteTemplate(w, "collections", data)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCollectionOfTheActualUser() service.Register {
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
		if register.Email == InfoOfUserWhoAreConnected.Email {
			return register
		}
	}
	return service.Register{}
}

func handleAddToCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req ArtifactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if IsLogin {
		// Appeler la fonction avec l'ID reçu
		AddToCollectionUser(req.Type, req.ID)

		// Réponse au client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Artifact added successfully",
			"id":      req.ID,
		})
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "You're not connected",
			"id":      "",
		})
	}
}

func AddToCollectionUser(Type string, id string) {
	if Type == "artifact" {
		artifact := service.GetAllDataAboutOneArtifact(id)
		temp := service.CollectionsStruct{
			Name:               artifact.Name,
			Image:              artifact.ImageURL,
			Type:               Type,
			LinkToTheRessource: "/artifacts/details?id=" + artifact.Id,
			DateAdded:          time.Now().Local().GoString(),
		}
		InfoOfUserWhoAreConnected.Collecttions.Collections = append(InfoOfUserWhoAreConnected.Collecttions.Collections, temp)
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
		if element.Email == InfoOfUserWhoAreConnected.Email && element.Password == InfoOfUserWhoAreConnected.Password {
			data[i].Collecttions = InfoOfUserWhoAreConnected.Collecttions
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

func RemoveCollections(w http.ResponseWriter, r *http.Request) {
	ElementToRemoveName, ElementToRemoveType := r.URL.Query().Get(("name")), r.URL.Query().Get(("type"))
	RemoveFromTheCollection(service.CollectionsStruct{
		Name: ElementToRemoveName,
		Type: ElementToRemoveType,
	})
	http.Redirect(w, r, "/collections", http.StatusSeeOther)
}

func RemoveFromTheCollection(ElementToRemove service.CollectionsStruct) {
	for i, element := range InfoOfUserWhoAreConnected.Collecttions.Collections {
		if element.Name == ElementToRemove.Name && element.Type == ElementToRemove.Type {
			InfoOfUserWhoAreConnected.Collecttions.Collections = append(InfoOfUserWhoAreConnected.Collecttions.Collections[:i], InfoOfUserWhoAreConnected.Collecttions.Collections[i+1:]...)
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
		if element.Email == InfoOfUserWhoAreConnected.Email && element.Password == InfoOfUserWhoAreConnected.Password {
			data[i].Collecttions = InfoOfUserWhoAreConnected.Collecttions
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
