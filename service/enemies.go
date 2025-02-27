package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetAllDetailsAboutOneEnemie(name string) (Enemies, error) {
	urlApi := "https://genshin.jmp.blue/enemies/" + name
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return Enemies{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return Enemies{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData Enemies
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			return Enemies{}, errDecode
		}
		image, err := checkResponseImageForEnemies(urlApi + "/list")
		if err != nil {
			log.Fatal(err)
		}
		if image {
			decodeData.ImageUrl = urlApi + "/icon"
		} else {
			decodeData.ImageUrl = "/static/image/NoImageAvaliable.webp"
		}
		return decodeData, nil
	}
	return Enemies{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetNamesOfAllTheEnemies() ([]string, error) {
	urlApi := "https://genshin.jmp.blue/enemies"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return []string{}, errReq
	}
	req.Header.Set("User-Agent", "Ynov Campus")
	res, Errres := httpClient.Do(req)
	if Errres != nil {
		return []string{}, Errres
	}
	if res.StatusCode == http.StatusOK {
		var decodeData []string
		errDecode := json.NewDecoder(res.Body).Decode(&decodeData)
		if errDecode != nil {
			fmt.Printf("error during decoding : %v\n", errDecode)
		}
		return decodeData, nil
	} else {
		fmt.Printf("Code error : %v, error message : %v", res.StatusCode, res.Status)
	}
	return []string{}, errors.New("Status code is not 200 but " + res.Status)
}

func GetAllEnemiesDetails() []Enemies {
	var err error
	var enemie Enemies
	AllNames, err := GetNamesOfAllTheEnemies()
	if err != nil {
		log.Fatal(err)
	}
	AllEnemiesDetail := []Enemies{}
	for _, name := range AllNames {
		if name == "cicin" {
			enemie = Enemies{
				Id: "cicin",
				Name: "Cicin",
				Region: "Multiple",
				Description: "N/A",
				Type: "Common Enemies",
				Family: "Mystical Beasts",
				Elements: []string{"Cryo", "Electro", "Hydro"},
				Drops: []EnemieDrop{},
				ElementalDescriptions: []EnemieElementalDescription{
					{
						Element: "Cryo",
						Description: "Little creatures that can ever so slightly manipulate Cryo. They and other such creatures are known collectively as Cicins. They are individually very weak, but it is precisely because they are weak that they have evolved unique methods of movement that allow them to avoid predators.",
					},  {
						Element: "Electro",
						Description: "Little creatures that can ever so slightly manipulate Electro. They and other such creatures are known collectively as Cicins, and though they are individually weak, they can cause significant damage under the right circumstances.",
					},  {
						Element: "Hydro",
						Description: "Little creatures that can ever so slightly manipulate Hydro. They and other such creatures are known collectively as Cicins, and they truly adore the rare and wonderful Mist Grass plant. It is by exploiting this relationship that the Fatui Mages have sought to commandeer Cicins in battle.",
					},
				},
				MoraGained: 0,
			}
		} else if name == "eye-of-the-storm" {
			enemie = Enemies{
				Id: "eye-of-the-storm",
				Name: "Eye of the Storm",
				Region: "Monstadt",
				Description: "N/A",
				Type: "Elite Enemies",
				Family: "Elemental Lifeforms",
				Elements: []string{"Cryo", "Pyro", "Hydro"},
				Drops: []EnemieDrop{},
				ElementalDescriptions: []EnemieElementalDescription{},
				MoraGained: 600,
			}
		} else if name == "the-great-snowboar-king" {
			enemie = Enemies{
				Id: "the-great-snowboar-king",
				Name: "The Great Snowboar King",
				Description: "N/A",
				Region: "Dragonspine",
				Type: "Elite Enemies",
				Family: "Unique Enemies",
				Elements: []string{},
				ElementalDescriptions: []EnemieElementalDescription{},
				Drops: []EnemieDrop{
					{
						Name: "Chilled Meat",
						Rarity: 1,
						MinimumLevel: 1,
					},
				},
				MoraGained: 0,
			}
		} else {
		enemie, err = GetAllDetailsAboutOneEnemie(name)
		if err != nil {
			log.Fatal(err)
		}}
		AllEnemiesDetail = append(AllEnemiesDetail, enemie)
	}
	return AllEnemiesDetail
}

func checkResponseImageForEnemies(apiURL string) (bool, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var responseArray []string
	if json.Unmarshal(body, &responseArray) == nil {
		if len(responseArray) == 2 && responseArray[0] == "icon" && responseArray[1] == "portrait" {
			return true, nil
		}
	}

	var responseMap map[string]string
	if json.Unmarshal(body, &responseMap) == nil {
		if val, exists := responseMap["error"]; exists && val == "No images for enemies/consecrated-beast exist" {
			return false, nil
		}
	}

	return false, nil
}
