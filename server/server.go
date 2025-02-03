package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

	"main.go/service"
)

var (
	funcMap = template.FuncMap{
		"add": add,
		"sub": sub,
	}
	TimeWhenConnect           time.Time     = time.Now()
	TimeToReload              time.Duration = 1 * time.Hour
	IsLoad                    bool          = false
	Mu                        sync.Mutex
	API_Data                  service.AllData
	Progress                  int
	Templates                 *template.Template
	IsLogin                   bool = false
	InfoOfUserWhoAreConnected service.Register
)

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

func InitServer() {
	var err error
	Templates, err = template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	//Initilalisation des routes
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/login/newregister", LoginNewRegister)
	http.HandleFunc("/login/register", LoginRegister)
	http.HandleFunc("/login/password-forgot", PasswordForgot)
	http.HandleFunc("/login/password-forgot/form", PasswordForgotData)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/collections", Collections)
	http.HandleFunc("/add-to-collection", handleAddToCollection)
	http.HandleFunc("/remove-from-the-collection", RemoveCollections)
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
	http.HandleFunc("/potions", Potions)
	http.HandleFunc("/potions/details", PotionsDetails)
	http.HandleFunc("/search", Search)
	http.HandleFunc("/loading", Loading)
	http.HandleFunc("/progress", ProgressF)
	http.HandleFunc("/profil", Profil)
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
	fmt.Println("http://localhost:8000/home")
	http.ListenAndServe("localhost:8000", nil)
}

func LoadAllData() {
	//food
	food := service.GetAllDetailsOfFood()
	Mu.Lock()
	API_Data.AllFood = food
	Progress = 10
	Mu.Unlock()
	//artifacts
	artifacts := service.GetAllArtifactsDetails()
	Mu.Lock()
	API_Data.ALLArtifacts = artifacts
	Progress = 20
	Mu.Unlock()
	//potions
	potions := service.GetAllPotions()
	Mu.Lock()
	API_Data.AllPotions = potions
	Progress = 30
	Mu.Unlock()
	//boss
	bosses := service.GetAllBossDetails()
	Mu.Lock()
	API_Data.AllBoss = bosses
	Progress = 40
	Mu.Unlock()
	//characters
	characters := service.GetAllCharactersDetails()
	Mu.Lock()
	API_Data.AllCharacters = characters
	Progress = 50
	Mu.Unlock()
	//Domain
	domains := service.GetAllDomainsDetails()
	Mu.Lock()
	API_Data.AllDomain = domains
	Progress = 60
	Mu.Unlock()
	//Elements
	elements := service.GetAllElementsDetails()
	Mu.Lock()
	API_Data.AllElement = elements
	Progress = 70
	Mu.Unlock()
	//Enemies
	enemies := service.GetAllEnemiesDetails()
	Mu.Lock()
	API_Data.AllEnnemies = enemies
	Progress = 80
	Mu.Unlock()
	//Weapons
	weapons := service.GetAllWeaponsDetails()
	Mu.Lock()
	API_Data.AllWeapons = weapons
	Progress = 90
	Mu.Unlock()

	Mu.Lock()
	Progress = 100
	Mu.Unlock()
}
