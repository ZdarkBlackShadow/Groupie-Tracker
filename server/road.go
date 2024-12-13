package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

func CharactersDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}
	fmt.Println(name)
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "%", "-")
	fmt.Println(name)
	err := Templates.ExecuteTemplate(w, "", service.GetAllDetailsAboutOneCharacters(name))
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func ArtifactsDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, "'", "-")
	name = strings.ReplaceAll(name, "%", "-")
	name = strings.ReplaceAll(name, " ", "-")
	Data := service.GetAllDataAboutOneArtifact(name)
	err := Templates.ExecuteTemplate(w, "ArtifactDetail", Data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
