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

type LoginData struct {
	IsLogin         bool
	IsWrongLogin    bool
	IsWrongRegister bool
}

var IsWrongLogin bool = false
var IsWrongRegister bool = false

func Login(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if IsLogin {
			http.Redirect(w, r, "/profil", http.StatusSeeOther)
		} else {
			Data := LoginData{
				IsLogin:         IsLogin,
				IsWrongLogin:    IsWrongLogin,
				IsWrongRegister: IsWrongRegister,
			}
			err1 := Templates.ExecuteTemplate(w, "login", Data)
			if err1 != nil {
				log.Fatal(err1)
			}
		}
	}
}

func LoginNewRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorCodeToSend.Update(401, "Unauthorized acces", "Acces forbiden", &ErrorToSend)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	} else {
		Email, Password, RetypedPassword := r.FormValue("NewRegisterEmail"), r.FormValue("NewRegisterPassword"), r.FormValue("NewRegisterRetypePassword")
		if Password != RetypedPassword {
			IsWrongRegister = true
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		DataToEncode := service.Register{
			Email:    utils.HashPassword(Email),
			Password: utils.HashPassword(Password),
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
			if element.Email == utils.HashPassword(Email) {
				AlreadyRegistered = true
				break
			}
		}
		if AlreadyRegistered {
			IsWrongRegister = true
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
			IsWrongRegister = false
			IsWrongLogin = false
			InfoOfUserWhoAreConnected = DataToEncode
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
}

func LoginRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorCodeToSend.Update(401, "Unauthorized acces", "Acces forbiden", &ErrorToSend)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
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
			if element.Email == utils.HashPassword(Email) {
				FindEmail = true
				if element.Password == utils.HashPassword(Password) {
					InfoOfUserWhoAreConnected = element
					IsLogin = true
					IsWrongLogin = false
					IsWrongRegister = false
					http.Redirect(w, r, "/home", http.StatusSeeOther)
				} else {
					IsWrongLogin = true
					http.Redirect(w, r, "/login", http.StatusSeeOther)
				}
				break
			}
		}
		if !FindEmail {
			IsWrongLogin = true
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

type PasswordForgotDataStruct struct {
	IsError bool
}

var IsErrorPasswordForgot bool = false

func PasswordForgot(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if IsLogin {
			http.Redirect(w, r, "/profil", http.StatusSeeOther)
		}
		Data := PasswordForgotDataStruct{
			IsError: IsErrorPasswordForgot,
		}
		IsErrorPasswordForgot = false
		err := Templates.ExecuteTemplate(w, "passwordforgot", Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func PasswordForgotData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorCodeToSend.Update(401, "Unauthorized acces", "Acces forbiden", &ErrorToSend)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	} else {
		Email, Password, RetypedPassword := r.FormValue("email"), r.FormValue("newpassword"), r.FormValue("newpassword2")
		if Password != RetypedPassword {
			IsErrorPasswordForgot = true
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
			if element.Email == utils.HashPassword(Email) {
				FindEmail = true
				fmt.Println(element.Password)
				data[i].Password = utils.HashPassword(Password)
				fmt.Println(element.Password)
				InfoOfUserWhoAreConnected = element
				break
			}
		}
		if !FindEmail {
			IsErrorPasswordForgot = true
			http.Redirect(w, r, "/login/password-forgot", http.StatusSeeOther)
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

func Logout(w http.ResponseWriter, r *http.Request) {
	IsLogin = false
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
