package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"main.go/service"
	"main.go/utils"
)

type DataCollectionsStruct struct {
	Data    service.Collecttions
	IsLogin bool
}

func Collections(w http.ResponseWriter, r *http.Request) {
	var data DataCollectionsStruct
	if IsLogin {
		data = DataCollectionsStruct{
			Data:    utils.GetCollectionOfTheActualUser(InfoOfUserWhoAreConnected).Collecttions,
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

func HandleAddToCollection(w http.ResponseWriter, r *http.Request) {
	if IsLogin {
		Type, id, detail := r.URL.Query().Get("type"), r.URL.Query().Get("id"), r.URL.Query().Get("detail")
		var temp service.CollectionsStruct
		var LinkToRedirect string
		var alreadyInCollection bool = false
		var err error
		if detail != "" {
			detail = "/details?id=" + id
		}
		switch Type {
		case "artifact":
			artifact, err := service.GetAllDataAboutOneArtifact(id)
			if err != nil {
				log.Fatal(err)
			}
			LinkToRedirect = "/artifacts" + detail
			temp = service.CollectionsStruct{
				Name:               artifact.Name,
				Image:              artifact.ImageURL,
				Type:               Type,
				LinkToTheRessource: "/artifacts/details?id=" + artifact.Id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "character":
			character, err := service.GetAllDetailsAboutOneCharacters(id)
			if err != nil {
				log.Fatal(err)
			}
			LinkToRedirect = "/characters" + detail
			temp = service.CollectionsStruct{
				Name:               character.Name,
				Image:              character.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/characters/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "domain":
			domain, err := service.GetAllDetailsAboutOneDomains(id)
			if err != nil {
				log.Fatal(err)
			}
			LinkToRedirect = "/domains" + detail
			temp = service.CollectionsStruct{
				Name:               domain.Name,
				Image:              "/static/image/Domain_Card.webp",
				Type:               Type,
				LinkToTheRessource: "/domains/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "element":
			element, err := service.GetAllDetailsAboutOneElements(id)
			if err != nil {
				log.Fatal(err)
			}
			LinkToRedirect = "/elements" + detail
			temp = service.CollectionsStruct{
				Name:               element.Name,
				Image:              element.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/elements/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "enemie":
			enemie, err := service.GetAllDetailsAboutOneEnemie(id)
			if err != nil {
				log.Fatal(err)
			}
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
			boss, err := service.GetAllDetailsAboutOneBoss(id)
			if err != nil {
				log.Fatal(err)
			}
			temp = service.CollectionsStruct{
				Name:               boss.Name,
				Image:              boss.ImageURL,
				Type:               Type,
				LinkToTheRessource: "/boss/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "food":
			LinkToRedirect = "/foods" + detail
			food := service.GetDetailsOfFood(id, AllDataOfAPI.AllFood)
			temp = service.CollectionsStruct{
				Name:               food.Name,
				Image:              food.ImageUrl,
				Type:               Type,
				LinkToTheRessource: "/foods/details?id=" + id,
				DateAdded:          time.Now().Local().GoString(),
			}
		case "potion":
			LinkToRedirect = "/potions" + detail
			potion := service.GetDetailsAboutOnePotion(id, AllDataOfAPI.AllPotions)
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
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func RemoveCollections(w http.ResponseWriter, r *http.Request) {
	ElementToRemoveName, ElementToRemoveType := r.URL.Query().Get(("name")), r.URL.Query().Get(("type"))
	utils.RemoveFromTheCollection(service.CollectionsStruct{
		Name: ElementToRemoveName,
		Type: ElementToRemoveType,
	}, &InfoOfUserWhoAreConnected)
	http.Redirect(w, r, "/collections", http.StatusSeeOther)
}
