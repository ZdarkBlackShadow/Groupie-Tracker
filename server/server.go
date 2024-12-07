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
	http.HandleFunc("/artifacts", Artifacts)
	http.HandleFunc("/boss", Boss)
	http.HandleFunc("/characters", Characters)
	http.HandleFunc("/domains", Domains)
	http.HandleFunc("/elements", Elements)
	http.HandleFunc("/enemies", Enemies)
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
	fileserver := http.FileServer(http.Dir("../assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Initialisation du serveur
	http.ListenAndServe("localhost:8000", nil)
}
