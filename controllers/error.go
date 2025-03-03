package controllers

import (
	"log"
	"net/http"
	"time"

	"main.go/utils"
)

var (
	ErrorCodeToSend utils.ErrorCodeStruct
	ErrorToSend     bool = false
)

func ErrorCode(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
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
}
