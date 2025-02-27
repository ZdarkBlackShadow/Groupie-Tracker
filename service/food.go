package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAllDetailsOfFood() (map[string]FoodStruct, error) {
	url := "https://genshin.jmp.blue/consumables/food"
	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	// Faire la requête GET
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Ynov Campus")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Vérifier le code de statut HTTP
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Code de statut différent de 200: " + res.Status)
	}

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body = append(body[:len(body)-13], body[len(body)-1:]...)
	var dishes map[string]FoodStruct
	if err := json.Unmarshal(body, &dishes); err != nil {
		// Affichage du message d'erreur si le décodage échoue
		fmt.Printf("Erreur lors du décodage : %v\n", err)
		return nil, err
	}
	if len(dishes) == 0 {
		return nil, errors.New("aucun plat trouvé dans la réponse")
	}
	return dishes, nil
}

func GetDetailsOfFood(id string, AllFood []FoodStruct) FoodStruct {
	for _, value := range AllFood {
		if value.Id == id {
			return value
		}
	}
	return FoodStruct{}
}

func TransformFoodToSlice(AllFood map[string]FoodStruct) []FoodStruct {
	var sliceFood []FoodStruct
	for name, value := range AllFood {
		if name != "id" {
			ToAppend := value
			ToAppend.ImageUrl = "https://genshin.jmp.blue/consumables/food/" + name
			ToAppend.Id = name
			sliceFood = append(sliceFood, ToAppend)
		}
	}
	return sliceFood
}
