package main

import (
	"RustyBoard/analytics"
	"RustyBoard/image_cacher"
	"RustyBoard/jira"
	"RustyBoard/persistence"
	"RustyBoard/server"
	"github.com/LastSprint/JiraGoIssues/services"
	"os"
)

const (
	PathToDbFileKey string = "RUSTY_BOARD_PATH_TO_DB_FILE"
	ServeAddress    string = "RUSTY_BOARD_SERVE_ADDRESS"
	JiraSearchUrl   string = "RUSTY_BOARD_JIRA_SEARCH_URL"
	JiraLogin       string = "RUSTY_BOARD_JIRA_LOGIN"
	JiraPass        string = "RUSTY_BOARD_JIRA_PASSWORD"
	StaticServerUrl string = "RUSTY_BOARD_STATIC_SERVER_URL"
)

func main() {
	srv := server.Api{
		DB:           createDb(),
		ServeAddr:    EnvOrCurrent(ServeAddress, "0.0.0.0:6644"),
		JiraAnalyzer: createJiraAnalyzer(),
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
				PathToFolderWithImages: "imgcache",
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
