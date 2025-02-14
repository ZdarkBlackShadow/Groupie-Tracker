package server

import (
	"log"
	"net/http"
)

func ErrorCode400(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "400", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode401(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "401", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode403(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "403", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode404(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "404", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode405(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "405", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode408(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "408", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode429(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "429", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode500(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "500", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode502(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "502", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode503(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "503", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ErrorCode504(w http.ResponseWriter, r *http.Request) {
	err := Templates.ExecuteTemplate(w, "504", nil)
	if err != nil {
		log.Fatal(err)
	}
}
