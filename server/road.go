package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"main.go/service"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err1 := Templates.ExecuteTemplate(w, "home", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	err1 := Templates.ExecuteTemplate(w, "login", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func LoginNewRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Acess refused")
	} else {
		Email, Password, RetypedPassword := r.FormValue("NewRegisterEmail"), r.FormValue("NewRegisterPassword"), r.FormValue("NewRegisterRetypePassword")
		if Password != RetypedPassword {
			fmt.Println("Password is not the same")
		}
		DataToEncode := Register{
			Email:    Email,
			Password: Password,
		}
		filePath := "./data/data.json"
		var data []Register
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
		}
		if len(fileContent) > 0 {
			if err := json.Unmarshal(fileContent, &data); err != nil {
				log.Fatalf("Erreur lors du décodage JSON : %v", err)
			}
		}
		AlreadyRegistered := false
		for _, element := range data {
			if element.Email == Email {
				AlreadyRegistered = true
				break
			}
		}
		if AlreadyRegistered {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			data = append(data, DataToEncode)
			updatedJsonData, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				log.Fatalf("Erreur lors de l'encodage JSON : %v", err)
			}
			if err := ioutil.WriteFile(filePath, updatedJsonData, 0644); err != nil {
				log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
			}
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
}

func LoginRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Acess refused")
	} else {
		Email, Password := r.FormValue("RegisterEmail"), r.FormValue("RegisterPassword")
		filePath := "./data/data.json"
		var data []Register
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture du fichier : %v", err)
		}
		if len(fileContent) > 0 {
			if err := json.Unmarshal(fileContent, &data); err != nil {
				log.Fatalf("Erreur lors du décodage JSON : %v", err)
			}
		}
		FindEmail := false
		for _, element := range data {
			if element.Email == Email {
				FindEmail = true
				if element.Password == Password {
					fmt.Println("Connexion réussie")
				} else {
					fmt.Println("wrong password")
				}
				break
			}
		}
		if !FindEmail {
			fmt.Println("Email not found")
		}
	}
}

func Artifacts(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.ArtifactDetails
	}
	Data := DataArtifacts{
		Data: service.GetAllArtifactsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "artifacts", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Boss(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.BossStruct
	}
	Data := DataArtifacts{
		Data: service.GetAllBossDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "boss", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Characters(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Characters
	}
	Data := DataArtifacts{
		Data: service.GetAllCharactersDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "characters", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Domains(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Domain
	}
	Data := DataArtifacts{
		Data: service.GetAllDomainsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "domains", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Elements(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Element
	}
	Data := DataArtifacts{
		Data: service.GetAllElementsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "elements", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Enemies(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Enemies
	}
	Data := DataArtifacts{
		Data: service.GetAllEnemiesDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "enemies", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func Weapons(w http.ResponseWriter, r *http.Request) {
	type DataArtifacts struct {
		Data []service.Weapon
	}
	Data := DataArtifacts{
		Data: service.GetAllWeaponsDetails(),
	}
	err1 := Templates.ExecuteTemplate(w, "weapon", Data)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func CharactersDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneCharacters(Id)
	err := Templates.ExecuteTemplate(w, "characterDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		fmt.Println(err)
	}
}

func ArtifactsDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDataAboutOneArtifact(Id)
	err := Templates.ExecuteTemplate(w, "ArtifactDetail", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func ElementDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneElements(Id)
	err := Templates.ExecuteTemplate(w, "ElementDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func BossDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneBoss(Id)
	err := Templates.ExecuteTemplate(w, "bossDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func DomainDetail(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneDomains(Id)
	err := Templates.ExecuteTemplate(w, "domainDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func EnemiesDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDetailsAboutOneEnemie(Id)
	err := Templates.ExecuteTemplate(w, "enemiesDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func WeaponDetails(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if Id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}
	Data := service.GetAllDataAboutOneWeapon(Id)
	err := Templates.ExecuteTemplate(w, "weaponsDetails", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func Food(w http.ResponseWriter, r *http.Request) {
	Data := service.GetAllDetailsOfFood()
	err := Templates.ExecuteTemplate(w, "food", Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
