package server

import (
	"log"
	"net/http"

	"main.go/service"
)

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
