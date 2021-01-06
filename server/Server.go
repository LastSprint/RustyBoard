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
	CacheTTL             time.Duration
	ImageCacheDirPath    string
	ImageCachePathPrefix string
	FrontendFilePath     string

	CertPath string
	KeyPath  string
}

func (api *Api) Run() {
	srv := http.NewServeMux()

	srv.HandleFunc("/projects", api.getAllProjects)

	fs := http.FileServer(http.Dir(api.ImageCacheDirPath))

	web := http.FileServer(http.Dir(api.FrontendFilePath))

	srv.Handle("/"+api.ImageCachePathPrefix+"/", http.StripPrefix("/"+api.ImageCachePathPrefix, fs))

	if len(api.FrontendFilePath) != 0 {
		srv.Handle("/", web)
	}

	handler := cors.Default().Handler(srv)

	log.Println("Server starts listening with params:")
	log.Println("CertPath:", api.CertPath)
	log.Println("KeyPath:", api.KeyPath)
	log.Println("ServeAddr:", api.ServeAddr)
	log.Println("CacheTTL:", api.CacheTTL)

	if len(api.CertPath) == 0 && len(api.KeyPath) == 0 {
		log.Fatal(http.ListenAndServe(api.ServeAddr, handler))
	} else {
		log.Fatal(http.ListenAndServeTLS(api.ServeAddr, api.CertPath, api.KeyPath, handler))
	}
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
