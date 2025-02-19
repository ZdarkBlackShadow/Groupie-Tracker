package controllers

import (
	"html/template"
	"time"

	"main.go/service"
	"main.go/utils"
)

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

type Register struct {
	Email       string
	Password    string
	Collections []struct {
		Name      string
		Image     string
		Type      string
		Link      string
		DateAdded string
	}
}

var (
	funcMap = template.FuncMap{
		"add": add,
		"sub": sub,
	}
	Templates                 *template.Template
	IsLogin                   bool = false
	IsLoad                    bool = false
	AllDataOfAPI              utils.AllData
	InfoOfUserWhoAreConnected service.Register
	TimeWhenConnect           time.Time     = time.Now()
	TimeToReload              time.Duration = 1 * time.Hour
)

func Init() error {
	var err error
	Templates, err = template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	return err
}
