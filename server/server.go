package server

import (
	"html/template"
	"net/http"
	"sync"
	"time"

	"main.go/service"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

var (
	funcMap = template.FuncMap{
		"add": add,
		"sub": sub,
	}
	TimeWhenConnect time.Time
	Mu              sync.Mutex
	API_Data        service.AllData
)

var Templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
var IsLoad bool = false

func InitServer() {
	//Initilalisation des routes
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/newregister", LoginNewRegister)
	http.HandleFunc("/login/register", LoginRegister)
	http.HandleFunc("/collections", Collections)
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
	http.HandleFunc("/search", Search)
	http.HandleFunc("/loading", Loading)
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

func TimeCheckerToReloadData() bool {
	/*
		Function who check if hour pass for reload the data,
		return true if it's more than one hour and false else
	*/
	if !IsLoad {
		API_Data = LoadAllData()
		IsLoad = true
		TimeWhenConnect = time.Now()
	}
	if time.Since(TimeWhenConnect) > time.Hour {
		API_Data = LoadAllData()
		TimeWhenConnect = time.Now()
		return true
	} else {
		return false
	}
}

func LoadAllData() service.AllData {
	result := service.AllData{
		ALLArtifacts:  service.GetAllArtifactsDetails(),
		AllBoss:       service.GetAllBossDetails(),
		AllCharacters: service.GetAllCharactersDetails(),
		AllDomain:     service.GetAllDomainsDetails(),
		AllElement:    service.GetAllElementsDetails(),
		AllEnnemies:   service.GetAllEnemiesDetails(),
		AllWeapons:    service.GetAllWeaponsDetails(),
	}
	return result
}
