package server

import (
	"log"
	"net/http"

	"main.go/service"
)

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
