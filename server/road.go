package server

import (
	"log"
	"net/http"

	"main.go/service"
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

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
