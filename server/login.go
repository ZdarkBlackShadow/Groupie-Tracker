package server

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
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
			Password: hashPassword(Password),
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
			IsLogin = true
			InfoOfUserWhoAreConnected = DataToEncode
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
					InfoOfUserWhoAreConnected = element
					IsLogin = true
					http.Redirect(w, r, "/home", http.StatusSeeOther)
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

func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkPasswordHash(password, hash string) bool {
	computedHash := hashPassword(password)
	// Comparaison sécurisée pour éviter les attaques par timing
	return subtle.ConstantTimeCompare([]byte(computedHash), []byte(hash)) == 1
}

func PasswordForgot(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "passwordforgot", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func PasswordForgotData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Acess refused")
	} else {
		Email, Password, RetypedPassword := r.FormValue("email"), r.FormValue("newpassword"), r.FormValue("newpassword2")
		if Password != RetypedPassword {
			http.Redirect(w, r, "/login/password-forgot", http.StatusSeeOther)
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
		FindEmail := false
		for i, element := range data {
			if element.Email == Email {
				FindEmail = true
				fmt.Println(element.Password)
				data[i].Password = Password
				fmt.Println(element.Password)
				InfoOfUserWhoAreConnected = element
				break
			}
		}
		if !FindEmail {
			fmt.Println("Email not found")
		} else {
			fmt.Println(data)
			updatedJsonData, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				log.Fatalf("Erreur lors de l'encodage JSON : %v", err)
			}
			if err := ioutil.WriteFile(filePath, updatedJsonData, 0644); err != nil {
				log.Fatalf("Erreur lors de l'écriture dans le fichier : %v", err)
			}
			IsLogin = true
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
}
