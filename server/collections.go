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
	if IsLogin {
		Type, id := r.URL.Query().Get("type"), r.URL.Query().Get("id")
		var temp service.CollectionsStruct
		var LinkToRedirect string
		switch Type {
		case "artifact":
			artifact := service.GetAllDataAboutOneArtifact(id)
			LinkToRedirect = "/artifacts"
			temp = service.CollectionsStruct{
				Name:               artifact.Name,
				Image:              artifact.ImageURL,
				Type:               Type,
				LinkToTheRessource: "/artifacts/details?id=" + artifact.Id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "character":
			character := service.GetAllDetailsAboutOneCharacters(id)
			LinkToRedirect = "/characters"
			temp = service.CollectionsStruct{
				Name:               character.Name,
				Image:              character.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "domain":
			domain := service.GetAllDetailsAboutOneDomains(id)
			LinkToRedirect = "/domains"
			temp = service.CollectionsStruct{
				Name:               domain.Name,
				Image:              domain.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "element":
			element := service.GetAllDetailsAboutOneElements(id)
			LinkToRedirect = "/elements"
			temp = service.CollectionsStruct{
				Name:               element.Name,
				Image:              element.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "enemie":
			enemie := service.GetAllDetailsAboutOneEnemie(id)
			LinkToRedirect = "/enemies"
			temp = service.CollectionsStruct{
				Name:               enemie.Name,
				Image:              enemie.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "weapon":
			LinkToRedirect = "/weapons"
			weapon := service.GetAllDataAboutOneWeapon(id)
			temp = service.CollectionsStruct{
				Name:               weapon.Name,
				Image:              weapon.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "boss":
			LinkToRedirect = "/boss"
			boss := service.GetAllDetailsAboutOneBoss(id)
			temp = service.CollectionsStruct{
				Name:               boss.Name,
				Image:              boss.ImageURL,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "food":
			LinkToRedirect = "/food"
			food := service.GetAllDetailsOfFood()
			temp = service.CollectionsStruct{
				Name:               food[id].Name,
				Image:              "",
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		case "potion":
			LinkToRedirect = "/potions"
			potion := service.GetDetailsAboutOnePotion(id, service.GetAllPotions())
			temp = service.CollectionsStruct{
				Name:               potion.Name,
				Image:              potion.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "",
				DateAdded:          time.Now().Local().GoString(),
			}
		default:
			fmt.Println("Error with the type")
		}
		InfoOfUserWhoAreConnected.Collecttions.Collections = append(InfoOfUserWhoAreConnected.Collecttions.Collections, temp)
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
		http.Redirect(w, r, LinkToRedirect, http.StatusSeeOther)
	} else {
		fmt.Println("You're not connected")
		http.Redirect(w, r, "/artifacts", http.StatusSeeOther)
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
