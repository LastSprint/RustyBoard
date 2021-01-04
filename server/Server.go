package server

import (
	"RustyBoard/jira"
	"RustyBoard/persistence"
	"RustyBoard/server/models"
	"encoding/json"
	"github.com/rs/cors"
	"log"
	"net/http"
	"sync"
	"time"
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
	CacheTTL time.Duration
	ImageCacheDirPath string
	ImageCachePathPrefix string
}

func (api *Api) Run() {
	srv := http.NewServeMux()

	srv.HandleFunc("/projects", api.getAllProjects)

	fs := http.FileServer(http.Dir(api.ImageCacheDirPath))

	srv.Handle("/"+api.ImageCachePathPrefix+"/", http.StripPrefix("/" + api.ImageCachePathPrefix, fs))

	handler := cors.Default().Handler(srv)

	http.ListenAndServe(api.ServeAddr, handler)
}

type ResponseCache struct {
	data models.GetAllProjectResponse
	time time.Time
}

var cache = ResponseCache{}

var mutex = &sync.Mutex{}

func (api *Api) getAllProjects(w http.ResponseWriter, r *http.Request) {

	if cache.time.Add(time.Minute * api.CacheTTL).After(time.Now()) {
		if err := json.NewEncoder(w).Encode(cache.data); err != nil {
			writeError(w, err, http.StatusInternalServerError)
		}
		return
	}


	data, err := api.DB.ReadAll()

	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	}

	analytics, errors := api.JiraAnalyzer.Run(data)

	result := models.GetAllProjectResponse{
		Data:   analytics,
		Errors: errors,
	}

	mutex.Lock()
	cache.time = time.Now()
	cache.data = result
	mutex.Unlock()
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
