package server

import (
	"log"
	"net/http"

	"main.go/service"
)

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
