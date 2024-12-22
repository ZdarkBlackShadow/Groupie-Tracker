package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var Templates, errTemplates = template.ParseGlob("./templates/*.html")

func InitServer() {
	if errTemplates != nil {
		//erreur lors de l'ouverture des templates
		fmt.Printf("Erreur => %s\n", errTemplates.Error())
		os.Exit(02)
	}
	//Initilalisation des routes
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/newregister", LoginNewRegister)
	http.HandleFunc("/login/register", LoginRegister)
	http.HandleFunc("/artifacts", Artifacts)
	http.HandleFunc("/artifacts/details", ArtifactsDetails)
	http.HandleFunc("/boss", Boss)
	http.HandleFunc("/boss/details", BossDetails)
	http.HandleFunc("/characters", Characters)
	http.HandleFunc("/characters/details", CharactersDetails)
	http.HandleFunc("/domains", Domains)
	http.HandleFunc("/domains/details", DomainDetail)
	http.HandleFunc("/elements", Elements)
	http.HandleFunc("/elements/details", ElementDetails)
	http.HandleFunc("/enemies", Enemies)
	http.HandleFunc("/enemies/details", EnemiesDetails)
	http.HandleFunc("/weapons", Weapons)
	http.HandleFunc("/weapons/details", WeaponDetails)
	http.HandleFunc("/food", Food)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/home":
			Home(w, r)
		case "/login":
			Login(w, r)
		default:
			http.Redirect(w, r, "/404", http.StatusAccepted)
		}
	})
	//Initialisation des assets
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
