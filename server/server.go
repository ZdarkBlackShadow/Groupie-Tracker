package server

import (
	"fmt"
	"log"
	"net/http"

	"main.go/controllers"
)

func InitServer() {
	var err error
	err = controllers.Init()
	if err != nil {
		log.Fatal(err)
	}
	//Initilalisation des routes
	http.HandleFunc("/home", controllers.Home)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/login/newregister", controllers.LoginNewRegister)
	http.HandleFunc("/login/register", controllers.LoginRegister)
	http.HandleFunc("/login/password-forgot", controllers.PasswordForgot)
	http.HandleFunc("/login/password-forgot/form", controllers.PasswordForgotData)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/collections", controllers.Collections)
	http.HandleFunc("/add-to-collection", controllers.HandleAddToCollection)
	http.HandleFunc("/remove-from-the-collection", controllers.RemoveCollections)
	http.HandleFunc("/artifacts", controllers.Artifacts)
	http.HandleFunc("/artifacts/details", controllers.ArtifactsDetails)
	http.HandleFunc("/boss", controllers.Boss)
	http.HandleFunc("/boss/details", controllers.BossDetails)
	http.HandleFunc("/characters", controllers.Characters)
	http.HandleFunc("/characters/details", controllers.CharactersDetails)
	http.HandleFunc("/domains", controllers.Domains)
	http.HandleFunc("/domains/details", controllers.DomainDetail)
	http.HandleFunc("/elements", controllers.Elements)
	http.HandleFunc("/elements/details", controllers.ElementDetails)
	http.HandleFunc("/enemies", controllers.Enemies)
	http.HandleFunc("/enemies/details", controllers.EnemiesDetails)
	http.HandleFunc("/weapons", controllers.Weapons)
	http.HandleFunc("/weapons/details", controllers.WeaponDetails)
	http.HandleFunc("/foods", controllers.Food)
	http.HandleFunc("/foods/details", controllers.FoodDetail)
	http.HandleFunc("/potions", controllers.Potions)
	http.HandleFunc("/potions/details", controllers.PotionsDetails)
	http.HandleFunc("/search", controllers.Search)
	http.HandleFunc("/loading", controllers.Loading)
	http.HandleFunc("/progress", controllers.ProgressF)
	http.HandleFunc("/profil", controllers.Profil)
	http.HandleFunc("/error", controllers.ErrorCode)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			http.Redirect(w, r, "/home", http.StatusFound)
		case "/home":
			controllers.Home(w, r)
		case "/login":
			controllers.Login(w, r)
		case "/login/newregister":
			controllers.LoginNewRegister(w, r)
		case "/login/register":
			controllers.LoginRegister(w, r)
		case "/login/password-forgot":
			controllers.PasswordForgot(w, r)
		case "/login/password-forgot/form":
			controllers.PasswordForgotData(w, r)
		case "/logout":
			controllers.Logout(w, r)
		case "/collections":
			controllers.Collections(w, r)
		case "/add-to-collection":
			controllers.HandleAddToCollection(w, r)
		case "/remove-from-the-collection":
			controllers.RemoveCollections(w, r)
		case "/artifacts":
			controllers.Artifacts(w, r)
		case "/artifacts/details":
			controllers.ArtifactsDetails(w, r)
		case "/boss":
			controllers.Boss(w, r)
		case "/boss/details":
			controllers.BossDetails(w, r)
		case "/characters":
			controllers.Characters(w, r)
		case "/characters/details":
			controllers.CharactersDetails(w, r)
		case "/domains":
			controllers.Domains(w, r)
		case "/domains/details":
			controllers.DomainDetail(w, r)
		case "/elements":
			controllers.Elements(w, r)
		case "/elements/details":
			controllers.ElementDetails(w, r)
		case "/enemies":
			controllers.Enemies(w, r)
		case "/enemies/details":
			controllers.EnemiesDetails(w, r)
		case "/weapons":
			controllers.Weapons(w, r)
		case "/weapons/details":
			controllers.WeaponDetails(w, r)
		case "/foods":
			controllers.Food(w, r)
		case "/foods/details":
			controllers.FoodDetail(w, r)
		case "/potions":
			controllers.Potions(w, r)
		case "/potions/details":
			controllers.PotionsDetails(w, r)
		case "/search":
			controllers.Search(w, r)
		case "/loading":
			controllers.Loading(w, r)
		case "/progress":
			controllers.ProgressF(w, r)
		case "/profil":
			controllers.Profil(w, r)
		case "/error":
			controllers.ErrorCode(w, r)
		default:
			controllers.ErrorCodeToSend.Update(404, "Not found", "Oops! The page you're looking for seems to have disappeared into Teyvat...", &controllers.ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		}
	})
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("http://localhost:8000/home")
	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
