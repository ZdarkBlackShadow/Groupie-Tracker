package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"main.go/service"
	"main.go/utils"
)

type DataArtifacts struct {
	Data        []service.ArtifactDetails
	TotalPages  int
	CurrentPage int
	IsLogin     bool
}

type DataArtifactDetails struct {
	Data    service.ArtifactDetails
	IsLogin bool
}

func Artifacts(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if len(r.URL.Query()) > 2 {
			ErrorCodeToSend.Update(400, "Bad Request", "Too much parameters in the URL", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			filters := utils.GetArtifactFilters(r)
			if !utils.ArtifactFilterChecker(filters) {
				ErrorCodeToSend.Update(400, "Bad Request", "Error in the parameter of the URL", &ErrorToSend)
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			} else {
				allArtifacts := utils.ApplyArtifactFilters(filters, AllDataOfAPI.ALLArtifacts)

				pageParam := r.URL.Query().Get("page")
				page, err := strconv.Atoi(pageParam)
				if err != nil || page < 1 {
					page = 1
				}
				pagedData, totalPages, page := utils.PaginateArtifacts(allArtifacts, page, 16)
				Data := DataArtifacts{
					Data:        pagedData,
					TotalPages:  totalPages,
					CurrentPage: page,
					IsLogin:     IsLogin,
				}
				err1 := Templates.ExecuteTemplate(w, "artifacts", Data)
				if err1 != nil {
					log.Fatal(err1)
				}
			}
		}
	}
}

func ArtifactsDetails(w http.ResponseWriter, r *http.Request) {
	if !IsLoad || time.Since(TimeWhenConnect) > TimeToReload {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		if len(r.URL.Query()) > 1 || !r.URL.Query().Has("id") {
			ErrorCodeToSend.Update(400, "Bad Request", "Too much parameters in the URL", &ErrorToSend)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
		} else {
			var err error
			Id := r.URL.Query().Get("id")
			artifact, err := service.GetAllDataAboutOneArtifact(Id)
			if err != nil {
				log.Fatal(err)
			}
			if artifact.Name == "" {
				http.Redirect(w, r, "/400", http.StatusSeeOther)
			} else {
				Data := DataArtifactDetails{
					Data:    artifact,
					IsLogin: IsLogin,
				}
				err := Templates.ExecuteTemplate(w, "ArtifactDetail", Data)
				if err != nil {
					http.Error(w, "Error rendering template", http.StatusInternalServerError)
				}
			}
		}
	}
}
