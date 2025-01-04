package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"main.go/service"
)

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
		DataToEncode := service.Register{
			Email:    Email,
			Password: Password,
		}
		filePath := "./data/data.json"
		var data []service.Register
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
		var data []service.Register
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
