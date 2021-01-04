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
	Run(data []persistence.StoredItem) ([]jira.ProjectData, []error)
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

	analytics, errors := api.JiraAnalyzer.Run(data)

	result := models.GetAllProjectResponse{
		Data:   analytics,
		Errors: errors,
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})

	if err != nil {
		log.Println("[WARN] Couldn't serialize error while writing json", err.Error())
	}
}
