package main

import (
	"RustyBoard/analytics"
	"RustyBoard/image_cacher"
	"RustyBoard/jira"
	"RustyBoard/persistence"
	"RustyBoard/server"
	"github.com/LastSprint/JiraGoIssues/services"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	PathToDbFileKey string = "RUSTY_BOARD_PATH_TO_DB_FILE"
	ServeAddress    string = "RUSTY_BOARD_SERVE_ADDRESS"
	JiraSearchUrl   string = "RUSTY_BOARD_JIRA_SEARCH_URL"
	JiraLogin       string = "RUSTY_BOARD_JIRA_LOGIN"
	JiraPass        string = "RUSTY_BOARD_JIRA_PASSWORD"
	StaticServerUrl string = "RUSTY_BOARD_STATIC_SERVER_URL"
	CacheTTL string = "RUSTY_BOARD_CACHE_TTL"
	ImageCacheDirPath string = "RUSTY_BOARD_IMAGE_CACHE_DIR_PATH"
	PathToCert string = "RUSTY_BOARD_TLS_CERT_PATH"
	PathToKey string = "RUSTY_BOARD_TLS_KEY_PATH"
	PathToFrontend string = "RUSTY_BOARD_PATH_TO_FRONTEND"
)

var imgCacheDirPath = EnvOrCurrent(ImageCacheDirPath, "imgcache")

func main() {
	srv := server.Api{
		DB:           createDb(),
		ServeAddr:    EnvOrCurrent(ServeAddress, "0.0.0.0:6644"),
		JiraAnalyzer: createJiraAnalyzer(),
		CacheTTL: readCacheTTL(),
		ImageCacheDirPath: imgCacheDirPath,
		ImageCachePathPrefix: imgCacheDirPath,
		CertPath: EnvOrCurrent(PathToCert, ""),
		KeyPath: EnvOrCurrent(PathToKey, ""),
		FrontendFilePath: EnvOrCurrent(PathToFrontend, ""),
	}

	srv.Run()
}

func createJiraAnalyzer() *analytics.JiraAnalytics {
	return &analytics.JiraAnalytics{
		JiraAnalyzer: &jira.Analyzer{
			&jira.JiraServiceWrapper{
				services.NewJiraIssueLoader(
					EnvOrCurrent(JiraSearchUrl, ""),
					EnvOrCurrent(JiraLogin, ""),
					EnvOrCurrent(JiraPass, ""),
				),
			},
			&image_cacher.AsyncImageCacher{
				PathToFolderWithImages: imgCacheDirPath,
				UrlPathToImages:        EnvOrCurrent(StaticServerUrl, "http://localhost:6644/imgcache"),
				User:                   EnvOrCurrent(JiraLogin, ""),
				Pass:                   EnvOrCurrent(JiraPass, ""),
			},
		},
	}
}

func createDb() *persistence.OneFileDB {
	db := &persistence.OneFileDB{
		PathToFile: EnvOrCurrent(PathToDbFileKey, "db.json"),
	}

	if err := db.Validate(); err != nil {
		panic(err)
	}

	return db
}

func EnvOrCurrent(key string, def string) string {

	env := os.Getenv(key)

	if len(env) == 0 {
		return def
	}

	return env
}

func readCacheTTL() time.Duration {
	ttlString := EnvOrCurrent(CacheTTL, "10")
	val, err := strconv.Atoi(ttlString)

	if err != nil {
		log.Println("[ERR] Couldn't parse CacheTTL", ttlString)
		return time.Duration(10)
	}

	return time.Duration(val)
}