package server

import (
	"RustyBoard/jira"
	"RustyBoard/persistence"
	"RustyBoard/server/models"
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type DB interface {
	ReadAll() ([]persistence.StoredItem, error)
}

type JiraAnalyzer interface {
	Analyze(jql []string, name string) (*jira.ProjectData, []error)
}

type Api struct {
	DB
	ServeAddr string
	JiraAnalyzer
}

func (api *Api) Run() {
	srv := http.NewServeMux()

	srv.HandleFunc("/projects", api.getAllProjects)

	handler := cors.Default().Handler(srv)

	http.ListenAndServe(api.ServeAddr, handler)
}

func (api *Api) getAllProjects(w http.ResponseWriter, r *http.Request) {

	data, err := api.DB.ReadAll()

	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}

	result := models.GetAllProjectResponse{}

	for _, item := range data {
		projectInfo, errs := api.JiraAnalyzer.Analyze(item.JiraQueries, item.ProjectName)

		for _, it := range errs {
			result.Errors = append(result.Errors, it)
		}

		if projectInfo == nil {
			continue
		}

		projectInfo.ImageUrl = item.ImageUrl

		result.Data = append(result.Data, *projectInfo)
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(map[string]string {"msg": err.Error()})

	if err != nil {
		log.Println("[WARN] Couldn't serialize error while writing json", err.Error())
	}
}