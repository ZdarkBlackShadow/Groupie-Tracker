package controllers

import (
	"log"
	"net/http"

	"main.go/utils"
)

var (
	ErrorCodeToSend utils.ErrorCodeStruct
	ErrorToSend     bool = false
)

func ErrorCode(w http.ResponseWriter, r *http.Request) {
	if !ErrorToSend {
		ErrorCodeToSend.Update(200, "Status Ok", "No errors detected into Teyvat", &ErrorToSend)
	}
	if len(r.URL.Query()) > 0 {
		ErrorCodeToSend.Update(400, "Bad Request", "Too much parameters in the URL", &ErrorToSend)
	}
	ErrorToSend = false
	err := Templates.ExecuteTemplate(w, "ErrorCode", ErrorCodeToSend)
	if err != nil {
		log.Fatal(err)
	}
}
