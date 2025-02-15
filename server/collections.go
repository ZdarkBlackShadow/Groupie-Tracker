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
		Type, id, detail := r.URL.Query().Get("type"), r.URL.Query().Get("id"), r.URL.Query().Get("detail")
		var temp service.CollectionsStruct
		var LinkToRedirect string
		var alreadyInCollection bool = false
		if detail != "" {
			detail = "/details?id=" + id
		}
		switch Type {
		case "artifact":
			artifact := service.GetAllDataAboutOneArtifact(id)
			LinkToRedirect = "/artifacts" + detail
			temp = service.CollectionsStruct{
				Name:               artifact.Name,
				Image:              artifact.ImageURL,
				Type:               Type,
				LinkToTheRessource: "/artifacts/details?id=" + artifact.Id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "character":
			character := service.GetAllDetailsAboutOneCharacters(id)
			LinkToRedirect = "/characters" + detail
			temp = service.CollectionsStruct{
				Name:               character.Name,
				Image:              character.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/characters/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "domain":
			domain := service.GetAllDetailsAboutOneDomains(id)
			LinkToRedirect = "/domains" + detail
			temp = service.CollectionsStruct{
				Name:               domain.Name,
				Image:              "/static/image/Domain_Card.webp",
				Type:               Type,
				LinkToTheRessource: "/domains/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "element":
			element := service.GetAllDetailsAboutOneElements(id)
			LinkToRedirect = "/elements" + detail
			temp = service.CollectionsStruct{
				Name:               element.Name,
				Image:              element.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/elements/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "enemie":
			enemie := service.GetAllDetailsAboutOneEnemie(id)
			LinkToRedirect = "/enemies" + detail
			temp = service.CollectionsStruct{
				Name:               enemie.Name,
				Image:              enemie.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/enemies/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "weapon":
			LinkToRedirect = "/weapons" + detail
			weapon := service.GetAllDataAboutOneWeapon(id)
			temp = service.CollectionsStruct{
				Name:               weapon.Name,
				Image:              weapon.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/weapons/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "boss":
			LinkToRedirect = "/boss"
			boss := service.GetAllDetailsAboutOneBoss(id)
			temp = service.CollectionsStruct{
				Name:               boss.Name,
				Image:              boss.ImageURL,
				Type:               Type,
				LinkToTheRessource: "/boss/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "food":
			LinkToRedirect = "/foods" + detail
			food := service.GetDetailsOfFood(id, API_Data.AllFood)
			temp = service.CollectionsStruct{
				Name:               food.Name,
				Image:              food.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/foods/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "potion":
			LinkToRedirect = "/potions" + detail
			potion := service.GetDetailsAboutOnePotion(id, service.GetAllPotions())
			temp = service.CollectionsStruct{
				Name:               potion.Name,
				Image:              potion.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/potions/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		default:
			fmt.Println("Error with the type")
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
				for _, element := range data[i].Collecttions.Collections {
					if element.Name == temp.Name && element.Type == temp.Type {
						fmt.Println("Element already in the collection")
						alreadyInCollection = true
						break
					}
				}
				if !alreadyInCollection {
					InfoOfUserWhoAreConnected.Collecttions.Collections = append(InfoOfUserWhoAreConnected.Collecttions.Collections, temp)
					data[i].Collecttions = InfoOfUserWhoAreConnected.Collecttions
				}
				break
			}
		}
		if !alreadyInCollection {
			updatedJsonData, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				log.Fatalf("Erreur lors de l'encodage JSON : %v", err)
			}
			if err := ioutil.WriteFile("./data/data.json", updatedJsonData, 0644); err != nil {
				log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
			}
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
